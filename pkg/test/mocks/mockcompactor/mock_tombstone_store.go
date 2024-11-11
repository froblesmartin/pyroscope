// Code generated by mockery. DO NOT EDIT.

package mockcompactor

import (
	bbolt "go.etcd.io/bbolt"

	iter "github.com/grafana/pyroscope/pkg/iter"

	metastorev1 "github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1"

	mock "github.com/stretchr/testify/mock"

	raft "github.com/hashicorp/raft"

	store "github.com/grafana/pyroscope/pkg/experiment/metastore/compaction/compactor/store"
)

// MockTombstoneStore is an autogenerated mock type for the TombstoneStore type
type MockTombstoneStore struct {
	mock.Mock
}

type MockTombstoneStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTombstoneStore) EXPECT() *MockTombstoneStore_Expecter {
	return &MockTombstoneStore_Expecter{mock: &_m.Mock}
}

// AddTombstones provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockTombstoneStore) AddTombstones(_a0 *bbolt.Tx, _a1 *raft.Log, _a2 *metastorev1.Tombstones) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for AddTombstones")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, *raft.Log, *metastorev1.Tombstones) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTombstoneStore_AddTombstones_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTombstones'
type MockTombstoneStore_AddTombstones_Call struct {
	*mock.Call
}

// AddTombstones is a helper method to define mock.On call
//   - _a0 *bbolt.Tx
//   - _a1 *raft.Log
//   - _a2 *metastorev1.Tombstones
func (_e *MockTombstoneStore_Expecter) AddTombstones(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockTombstoneStore_AddTombstones_Call {
	return &MockTombstoneStore_AddTombstones_Call{Call: _e.mock.On("AddTombstones", _a0, _a1, _a2)}
}

func (_c *MockTombstoneStore_AddTombstones_Call) Run(run func(_a0 *bbolt.Tx, _a1 *raft.Log, _a2 *metastorev1.Tombstones)) *MockTombstoneStore_AddTombstones_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*bbolt.Tx), args[1].(*raft.Log), args[2].(*metastorev1.Tombstones))
	})
	return _c
}

func (_c *MockTombstoneStore_AddTombstones_Call) Return(_a0 error) *MockTombstoneStore_AddTombstones_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTombstoneStore_AddTombstones_Call) RunAndReturn(run func(*bbolt.Tx, *raft.Log, *metastorev1.Tombstones) error) *MockTombstoneStore_AddTombstones_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTombstones provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockTombstoneStore) DeleteTombstones(_a0 *bbolt.Tx, _a1 *raft.Log, _a2 *metastorev1.Tombstones) error {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTombstones")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, *raft.Log, *metastorev1.Tombstones) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTombstoneStore_DeleteTombstones_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTombstones'
type MockTombstoneStore_DeleteTombstones_Call struct {
	*mock.Call
}

// DeleteTombstones is a helper method to define mock.On call
//   - _a0 *bbolt.Tx
//   - _a1 *raft.Log
//   - _a2 *metastorev1.Tombstones
func (_e *MockTombstoneStore_Expecter) DeleteTombstones(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockTombstoneStore_DeleteTombstones_Call {
	return &MockTombstoneStore_DeleteTombstones_Call{Call: _e.mock.On("DeleteTombstones", _a0, _a1, _a2)}
}

func (_c *MockTombstoneStore_DeleteTombstones_Call) Run(run func(_a0 *bbolt.Tx, _a1 *raft.Log, _a2 *metastorev1.Tombstones)) *MockTombstoneStore_DeleteTombstones_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*bbolt.Tx), args[1].(*raft.Log), args[2].(*metastorev1.Tombstones))
	})
	return _c
}

func (_c *MockTombstoneStore_DeleteTombstones_Call) Return(_a0 error) *MockTombstoneStore_DeleteTombstones_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTombstoneStore_DeleteTombstones_Call) RunAndReturn(run func(*bbolt.Tx, *raft.Log, *metastorev1.Tombstones) error) *MockTombstoneStore_DeleteTombstones_Call {
	_c.Call.Return(run)
	return _c
}

// Exists provides a mock function with given fields: _a0
func (_m *MockTombstoneStore) Exists(_a0 *metastorev1.BlockMeta) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Exists")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(*metastorev1.BlockMeta) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockTombstoneStore_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type MockTombstoneStore_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - _a0 *metastorev1.BlockMeta
func (_e *MockTombstoneStore_Expecter) Exists(_a0 interface{}) *MockTombstoneStore_Exists_Call {
	return &MockTombstoneStore_Exists_Call{Call: _e.mock.On("Exists", _a0)}
}

func (_c *MockTombstoneStore_Exists_Call) Run(run func(_a0 *metastorev1.BlockMeta)) *MockTombstoneStore_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*metastorev1.BlockMeta))
	})
	return _c
}

func (_c *MockTombstoneStore_Exists_Call) Return(_a0 bool) *MockTombstoneStore_Exists_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTombstoneStore_Exists_Call) RunAndReturn(run func(*metastorev1.BlockMeta) bool) *MockTombstoneStore_Exists_Call {
	_c.Call.Return(run)
	return _c
}

// GetTombstones provides a mock function with given fields: _a0, _a1
func (_m *MockTombstoneStore) GetTombstones(_a0 *bbolt.Tx, _a1 *raft.Log) (*metastorev1.Tombstones, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetTombstones")
	}

	var r0 *metastorev1.Tombstones
	var r1 error
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, *raft.Log) (*metastorev1.Tombstones, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(*bbolt.Tx, *raft.Log) *metastorev1.Tombstones); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metastorev1.Tombstones)
		}
	}

	if rf, ok := ret.Get(1).(func(*bbolt.Tx, *raft.Log) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTombstoneStore_GetTombstones_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTombstones'
type MockTombstoneStore_GetTombstones_Call struct {
	*mock.Call
}

// GetTombstones is a helper method to define mock.On call
//   - _a0 *bbolt.Tx
//   - _a1 *raft.Log
func (_e *MockTombstoneStore_Expecter) GetTombstones(_a0 interface{}, _a1 interface{}) *MockTombstoneStore_GetTombstones_Call {
	return &MockTombstoneStore_GetTombstones_Call{Call: _e.mock.On("GetTombstones", _a0, _a1)}
}

func (_c *MockTombstoneStore_GetTombstones_Call) Run(run func(_a0 *bbolt.Tx, _a1 *raft.Log)) *MockTombstoneStore_GetTombstones_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*bbolt.Tx), args[1].(*raft.Log))
	})
	return _c
}

func (_c *MockTombstoneStore_GetTombstones_Call) Return(_a0 *metastorev1.Tombstones, _a1 error) *MockTombstoneStore_GetTombstones_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTombstoneStore_GetTombstones_Call) RunAndReturn(run func(*bbolt.Tx, *raft.Log) (*metastorev1.Tombstones, error)) *MockTombstoneStore_GetTombstones_Call {
	_c.Call.Return(run)
	return _c
}

// ListEntries provides a mock function with given fields: _a0
func (_m *MockTombstoneStore) ListEntries(_a0 *bbolt.Tx) iter.Iterator[store.BlockEntry] {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ListEntries")
	}

	var r0 iter.Iterator[store.BlockEntry]
	if rf, ok := ret.Get(0).(func(*bbolt.Tx) iter.Iterator[store.BlockEntry]); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iter.Iterator[store.BlockEntry])
		}
	}

	return r0
}

// MockTombstoneStore_ListEntries_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListEntries'
type MockTombstoneStore_ListEntries_Call struct {
	*mock.Call
}

// ListEntries is a helper method to define mock.On call
//   - _a0 *bbolt.Tx
func (_e *MockTombstoneStore_Expecter) ListEntries(_a0 interface{}) *MockTombstoneStore_ListEntries_Call {
	return &MockTombstoneStore_ListEntries_Call{Call: _e.mock.On("ListEntries", _a0)}
}

func (_c *MockTombstoneStore_ListEntries_Call) Run(run func(_a0 *bbolt.Tx)) *MockTombstoneStore_ListEntries_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*bbolt.Tx))
	})
	return _c
}

func (_c *MockTombstoneStore_ListEntries_Call) Return(_a0 iter.Iterator[store.BlockEntry]) *MockTombstoneStore_ListEntries_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTombstoneStore_ListEntries_Call) RunAndReturn(run func(*bbolt.Tx) iter.Iterator[store.BlockEntry]) *MockTombstoneStore_ListEntries_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTombstoneStore creates a new instance of MockTombstoneStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTombstoneStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTombstoneStore {
	mock := &MockTombstoneStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
