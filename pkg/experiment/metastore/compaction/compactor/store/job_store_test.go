package store

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.etcd.io/bbolt"

	"github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1/raft_log"
)

func Test_JobStateStore(t *testing.T) {
	tempDir := t.TempDir()
	db, err := bbolt.Open(filepath.Join(tempDir, "boltdb"), 0644, nil)
	require.NoError(t, err)

	s := NewJobStore()
	tx, err := db.Begin(true)
	require.NoError(t, err)
	require.NoError(t, s.CreateBuckets(tx))
	assert.NoError(t, s.StoreJobState(tx, &raft_log.CompactionJobState{Name: "1"}))
	assert.NoError(t, s.StoreJobState(tx, &raft_log.CompactionJobState{Name: "2"}))
	assert.NoError(t, s.StoreJobState(tx, &raft_log.CompactionJobState{Name: "3"}))
	require.NoError(t, tx.Commit())

	s = NewJobStore()
	tx, err = db.Begin(true)
	require.NoError(t, err)
	state, err := s.GetJobState(tx, "2")
	require.NoError(t, err)
	assert.Equal(t, "2", state.Name)
	require.NoError(t, s.DeleteJobState(tx, "2"))
	state, err = s.GetJobState(tx, "2")
	require.ErrorIs(t, err, ErrorNotFound)
	require.Nil(t, state)
	require.NoError(t, tx.Commit())

	tx, err = db.Begin(true)
	require.NoError(t, err)

	iter := s.ListEntries(tx)
	expected := []string{"1", "3"}
	var i int
	for iter.Next() {
		assert.Equal(t, expected[i], iter.At().Name)
		i++
	}
	assert.Nil(t, iter.Err())
	assert.Nil(t, iter.Close())
	require.NoError(t, tx.Rollback())
}

func Test_JobPlanStore(t *testing.T) {
	tempDir := t.TempDir()
	db, err := bbolt.Open(filepath.Join(tempDir, "boltdb"), 0644, nil)
	require.NoError(t, err)

	s := NewJobStore()
	tx, err := db.Begin(true)
	require.NoError(t, err)
	require.NoError(t, s.CreateBuckets(tx))
	assert.NoError(t, s.StoreJobPlan(tx, &raft_log.CompactionJobPlan{Name: "1"}))
	require.NoError(t, tx.Commit())

	s = NewJobStore()
	tx, err = db.Begin(false)
	require.NoError(t, err)
	state, err := s.GetJobPlan(tx, "2")
	require.ErrorIs(t, err, ErrorNotFound)
	require.Nil(t, state)
	state, err = s.GetJobPlan(tx, "1")
	require.NoError(t, err)
	assert.Equal(t, "1", state.Name)
	require.NoError(t, tx.Rollback())

	tx, err = db.Begin(true)
	require.NoError(t, err)
	require.NoError(t, s.DeleteJobPlan(tx, "1"))
	require.NoError(t, tx.Commit())

	tx, err = db.Begin(false)
	require.NoError(t, err)
	state, err = s.GetJobPlan(tx, "1")
	require.ErrorIs(t, err, ErrorNotFound)
	require.Nil(t, state)
	require.NoError(t, tx.Rollback())
}
