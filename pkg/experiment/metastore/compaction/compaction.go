package compaction

import (
	"github.com/hashicorp/raft"
	"go.etcd.io/bbolt"

	metastorev1 "github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1"
	"github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1/raft_log"
)

type Compactor interface {
	// AddBlock enqueues a new block for compaction.
	// Implementation: the method must be idempotent.
	AddBlock(*bbolt.Tx, *raft.Log, *metastorev1.BlockMeta) error
}

type Planner interface {
	// NewPlan is used to plan new jobs. The proposed changes will then be
	// submitted for Raft consensus, with the leader's jobs being accepted
	// as the final decision.
	// Implementation: Plan must not change the state of Planner.
	NewPlan(*bbolt.Tx, *raft.Log) Plan

	// Scheduled must be called for each job after it is scheduled
	// to remove the job from future plans.
	// Implementation: the method must be idempotent.
	Scheduled(*bbolt.Tx, ...*raft_log.CompactionJobUpdate) error
}

type Plan interface {
	CreateJob() (*raft_log.CompactionJobPlan, error)
}

type Scheduler interface {
	// NewSchedule is used to plan a schedule update. The proposed schedule
	// will then be submitted for Raft consensus, with the leader's schedule
	// being accepted as the final decision.
	// Implementation: Schedule must not change the state of Scheduler.
	NewSchedule(*bbolt.Tx, *raft.Log) Schedule

	// UpdateSchedule adds new jobs and updates state of existing ones.
	// Implementation: the method must be idempotent.
	UpdateSchedule(*bbolt.Tx, *raft_log.CompactionPlanUpdate) error
}

type Schedule interface {
	// UpdateJob is called on behalf of the worker to update the job status.
	// A nil response should be interpreted as "no new lease": stop the work.
	// The scheduler must validate that the worker is allowed to update the job,
	// by comparing the fencing token of the job. Refer to the documentation for
	// details.
	UpdateJob(*metastorev1.CompactionJobStatusUpdate) (*raft_log.CompactionJobUpdate, error)

	// AssignJob is called on behalf of the worker to request a new job.
	AssignJob() (*raft_log.CompactionJobUpdate, error)

	// AddJob is called on behalf of the planner to add a new job to the schedule.
	AddJob(*raft_log.CompactionJobPlan) (*raft_log.CompactionJobUpdate, error)
}
