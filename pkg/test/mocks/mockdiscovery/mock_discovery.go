// Code generated by mockery. DO NOT EDIT.

package mockdiscovery

import (
	discovery "github.com/grafana/pyroscope/pkg/experiment/metastore/discovery"
	mock "github.com/stretchr/testify/mock"
)

// MockDiscovery is an autogenerated mock type for the Discovery type
type MockDiscovery struct {
	mock.Mock
}

type MockDiscovery_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDiscovery) EXPECT() *MockDiscovery_Expecter {
	return &MockDiscovery_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockDiscovery) Close() {
	_m.Called()
}

// MockDiscovery_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockDiscovery_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockDiscovery_Expecter) Close() *MockDiscovery_Close_Call {
	return &MockDiscovery_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockDiscovery_Close_Call) Run(run func()) *MockDiscovery_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDiscovery_Close_Call) Return() *MockDiscovery_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockDiscovery_Close_Call) RunAndReturn(run func()) *MockDiscovery_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Rediscover provides a mock function with given fields:
func (_m *MockDiscovery) Rediscover() {
	_m.Called()
}

// MockDiscovery_Rediscover_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rediscover'
type MockDiscovery_Rediscover_Call struct {
	*mock.Call
}

// Rediscover is a helper method to define mock.On call
func (_e *MockDiscovery_Expecter) Rediscover() *MockDiscovery_Rediscover_Call {
	return &MockDiscovery_Rediscover_Call{Call: _e.mock.On("Rediscover")}
}

func (_c *MockDiscovery_Rediscover_Call) Run(run func()) *MockDiscovery_Rediscover_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDiscovery_Rediscover_Call) Return() *MockDiscovery_Rediscover_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockDiscovery_Rediscover_Call) RunAndReturn(run func()) *MockDiscovery_Rediscover_Call {
	_c.Call.Return(run)
	return _c
}

// Subscribe provides a mock function with given fields: updates
func (_m *MockDiscovery) Subscribe(updates discovery.Updates) {
	_m.Called(updates)
}

// MockDiscovery_Subscribe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subscribe'
type MockDiscovery_Subscribe_Call struct {
	*mock.Call
}

// Subscribe is a helper method to define mock.On call
//   - updates discovery.Updates
func (_e *MockDiscovery_Expecter) Subscribe(updates interface{}) *MockDiscovery_Subscribe_Call {
	return &MockDiscovery_Subscribe_Call{Call: _e.mock.On("Subscribe", updates)}
}

func (_c *MockDiscovery_Subscribe_Call) Run(run func(updates discovery.Updates)) *MockDiscovery_Subscribe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(discovery.Updates))
	})
	return _c
}

func (_c *MockDiscovery_Subscribe_Call) Return() *MockDiscovery_Subscribe_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockDiscovery_Subscribe_Call) RunAndReturn(run func(discovery.Updates)) *MockDiscovery_Subscribe_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDiscovery creates a new instance of MockDiscovery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDiscovery(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDiscovery {
	mock := &MockDiscovery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
