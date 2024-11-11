package store

import (
	"go.etcd.io/bbolt"

	"github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1/raft_log"
	"github.com/grafana/pyroscope/pkg/iter"
)

type JobStore struct {
	JobStateStore
	JobPlanStore
}

var (
	jobStateBucketName = []byte("compaction_job_state")
	jobPlanBucketName  = []byte("compaction_job_plan")
)

func NewJobStore() *JobStore {
	return &JobStore{
		JobStateStore: JobStateStore{bucketName: jobStateBucketName},
		JobPlanStore:  JobPlanStore{bucketName: jobPlanBucketName},
	}
}

func (s JobStore) CreateBuckets(tx *bbolt.Tx) error {
	if _, err := tx.CreateBucketIfNotExists(jobStateBucketName); err != nil {
		return err
	}
	if _, err := tx.CreateBucketIfNotExists(jobPlanBucketName); err != nil {
		return err
	}
	return nil
}

type JobPlanStore struct{ bucketName []byte }

func (s JobPlanStore) StoreJobPlan(tx *bbolt.Tx, plan *raft_log.CompactionJobPlan) error {
	v, err := plan.MarshalVT()
	if err != nil {
		return err
	}
	return tx.Bucket(s.bucketName).Put([]byte(plan.Name), v)
}

func (s JobPlanStore) GetJobPlan(tx *bbolt.Tx, name string) (*raft_log.CompactionJobPlan, error) {
	b := tx.Bucket(s.bucketName).Get([]byte(name))
	if b == nil {
		return nil, ErrorNotFound
	}
	var v raft_log.CompactionJobPlan
	if err := v.UnmarshalVT(b); err != nil {
		return nil, err
	}
	return &v, nil
}

func (s JobPlanStore) DeleteJobPlan(tx *bbolt.Tx, name string) error {
	return tx.Bucket(s.bucketName).Delete([]byte(name))
}

type JobStateStore struct{ bucketName []byte }

func (s JobStateStore) bucket(tx *bbolt.Tx) *bbolt.Bucket { return tx.Bucket(s.bucketName) }

func (s JobStateStore) GetJobState(tx *bbolt.Tx, name string) (*raft_log.CompactionJobState, error) {
	b := s.bucket(tx).Get([]byte(name))
	if b == nil {
		return nil, ErrorNotFound
	}
	var v raft_log.CompactionJobState
	if err := v.UnmarshalVT(b); err != nil {
		return nil, err
	}
	return &v, nil
}

func (s JobStateStore) StoreJobState(tx *bbolt.Tx, state *raft_log.CompactionJobState) error {
	v, err := state.MarshalVT()
	if err != nil {
		return err
	}
	return s.bucket(tx).Put([]byte(state.Name), v)
}

func (s JobStateStore) DeleteJobState(tx *bbolt.Tx, name string) error {
	return s.bucket(tx).Delete([]byte(name))
}

func (s JobStateStore) ListEntries(tx *bbolt.Tx) iter.Iterator[*raft_log.CompactionJobState] {
	return newJobEntriesIterator(s.bucket(tx))
}

type jobEntriesIterator struct {
	iter *cursorIterator
	cur  *raft_log.CompactionJobState
	err  error
}

func newJobEntriesIterator(bucket *bbolt.Bucket) *jobEntriesIterator {
	return &jobEntriesIterator{iter: newCursorIter(nil, bucket.Cursor())}
}

func (x *jobEntriesIterator) Next() bool {
	if x.err != nil || !x.iter.Next() {
		return false
	}
	e := x.iter.At()
	if e.value == nil {
		x.err = ErrorNotFound
		return false
	}
	var s raft_log.CompactionJobState
	x.err = s.UnmarshalVT(e.value)
	if x.err != nil {
		return false
	}
	x.cur = &s
	return true
}

func (x *jobEntriesIterator) At() *raft_log.CompactionJobState { return x.cur }

func (x *jobEntriesIterator) Close() error { return x.iter.Close() }

func (x *jobEntriesIterator) Err() error {
	if err := x.iter.Err(); err != nil {
		return err
	}
	return x.err
}
