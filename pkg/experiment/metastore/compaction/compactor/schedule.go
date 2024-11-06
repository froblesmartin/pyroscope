package compactor

import (
	"container/heap"
	"slices"
	"time"

	"github.com/hashicorp/raft"
	"go.etcd.io/bbolt"

	metastorev1 "github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1"
	"github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1/raft_log"
)

type compactionSchedule struct {
	tx        *bbolt.Tx
	raft      *raft.Log
	scheduler *Scheduler
	assigner  *jobAssigner
}

func (p *compactionSchedule) UpdateJob(update *metastorev1.CompactionJobStatusUpdate) (*raft_log.CompactionJobState, error) {
	state, err := p.scheduler.store.GetJobState(p.tx, update.Name)
	if err != nil {
		return nil, err
	}
	if state == nil {
		// The job is not found. This may happen if the job has been
		// reassigned and completed by another worker.
		return nil, nil
	}
	if state.Token > update.Token {
		// The job is not assigned to this worker.
		return nil, nil
	}

	switch update.Status {
	default:
		// Not allowed and unknown status updates are ignored: eventually,
		// the job will be reassigned. The same for status handlers: a nil
		// state is returned, which is interpreted as "no new lease":
		// stop the work.
		return nil, nil

	case metastorev1.CompactionJobStatus_COMPACTION_STATUS_IN_PROGRESS:
		return p.handleInProgress(state), nil
	case metastorev1.CompactionJobStatus_COMPACTION_STATUS_SUCCESS:
		return p.handleSuccess(state, update)
	case metastorev1.CompactionJobStatus_COMPACTION_STATUS_FAILURE:
		return p.handleFailure(state), nil
	}
}

func (p *compactionSchedule) handleSuccess(
	state *raft_log.CompactionJobState,
	update *metastorev1.CompactionJobStatusUpdate,
) (*raft_log.CompactionJobState, error) {
	prepared, err := p.scheduler.store.GetCompactedBlocks(p.tx, state.Name)
	if err != nil {
		return nil, err
	}
	compacted := prepared.CloneVT()
	compacted.CompactedBlocks = update.CompactedBlocks
	compacted.DeletedBlocks = update.DeletedBlocks
	updated := state.CloneVT()
	updated.CompactedBlocks = compacted
	return updated, nil
}

func (p *compactionSchedule) handleInProgress(state *raft_log.CompactionJobState) *raft_log.CompactionJobState {
	updated := state.CloneVT()
	updated.LeaseExpiresAt = p.assigner.allocateLease()
	return updated
}

func (p *compactionSchedule) handleFailure(state *raft_log.CompactionJobState) *raft_log.CompactionJobState {
	updated := state.CloneVT()
	// The worker shouldn't be able to take the job back immediately, using
	// the current raft index as the token at assignment. Jobs are assigned
	// before status updates, however we revoke the job explicitly.
	updated.Token = p.raft.Index + 1
	// Zero lease: the job will be prioritized for reassignment.
	updated.LeaseExpiresAt = 0
	updated.Status = metastorev1.CompactionJobStatus_COMPACTION_STATUS_IN_PROGRESS
	updated.Failures++
	if p.scheduler.config.MaxFailures > 0 && updated.Failures >= p.scheduler.config.MaxFailures {
		updated.Status = metastorev1.CompactionJobStatus_COMPACTION_STATUS_CANCELLED
	}
	return updated
}

func (p *compactionSchedule) AssignJob() (*metastorev1.CompactionJob, *raft_log.CompactionJobState, error) {
	state := p.assigner.assign()
	if state == nil {
		return nil, nil, nil
	}
	job, err := p.scheduler.store.GetJob(p.tx, state.Name)
	if err != nil {
		return nil, nil, err
	}
	if job == nil {
		// Job not found. This should never happen and likely indicates
		// a data inconsistency. If we keep the job in the queue (as it
		// cannot be assigned), it will be dangling there forever.
		// Therefore, we remove it now: this is an exceptional case –
		// no state should be changed in compactionSchedule.
		p.deleteDangling(state)
		return nil, nil, nil
	}
	return job, state, nil
}

func (p *compactionSchedule) deleteDangling(state *raft_log.CompactionJobState) {
	_ = p.scheduler.store.DeleteJobState(p.tx, state.Name)
	_ = p.scheduler.store.DeleteJob(p.tx, state.Name)
	p.assigner.queue.delete(state)
}

type jobAssigner struct {
	raft   *raft.Log
	lease  time.Duration
	queue  *jobQueue
	copied []priorityQueue
	level  int
}

func (a *jobAssigner) assign() *raft_log.CompactionJobState {
	// We don't need to check the job ownership here: the worker asks
	// for a job assigment (new ownership).

	for a.level < len(a.queue.levels) {
		pq := a.queueLevelCopy(a.level)
		if pq.Len() == 0 {
			a.level++
			continue
		}

		switch job := heap.Pop(pq).(*jobEntry); job.Status {
		case metastorev1.CompactionJobStatus_COMPACTION_STATUS_UNSPECIFIED:
			return a.assignJob(job)

		case metastorev1.CompactionJobStatus_COMPACTION_STATUS_IN_PROGRESS:
			if a.now().UnixNano() > job.LeaseExpiresAt {
				return a.reassignJob(job)
			}
		}

		a.level++
	}

	return nil
}

func (a *jobAssigner) now() time.Time { return a.raft.AppendedAt }

func (a *jobAssigner) allocateLease() int64 { return a.now().Add(a.lease).UnixNano() }

func (a *jobAssigner) assignJob(e *jobEntry) *raft_log.CompactionJobState {
	job := e.CompactionJobState.CloneVT()
	job.Status = metastorev1.CompactionJobStatus_COMPACTION_STATUS_IN_PROGRESS
	job.LeaseExpiresAt = a.allocateLease()
	job.Token = a.raft.Index
	return job
}

func (a *jobAssigner) reassignJob(e *jobEntry) *raft_log.CompactionJobState {
	return a.assignJob(e)
}

// The queue must not be modified by assigner. Therefore, we're copying the
// queue levels lazily. The queue is supposed to be small (dozens of jobs
// running concurrently); in the worst case, we have a ~24b alloc per entry.
// Alternatively, we could push back the jobs to the queue, but it would
// require an explicit rollback call.
func (a *jobAssigner) queueLevelCopy(i int) *priorityQueue {
	s := i + 1 // Levels are 0-based.
	if s >= len(a.copied) || len(a.copied[i]) == 0 {
		a.copied = slices.Grow(a.copied, s)[:s]
		level := *a.queue.level(uint32(i))
		a.copied[i] = make(priorityQueue, len(level))
		for j, job := range level {
			jobCopy := *job
			a.copied[i][j] = &jobCopy
		}
	}
	return &a.copied[i]
}
