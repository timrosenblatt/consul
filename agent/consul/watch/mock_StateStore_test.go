// Code generated by mockery v2.12.2. DO NOT EDIT.

package watch

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockStateStore is an autogenerated mock type for the StateStore type
type MockStateStore struct {
	mock.Mock
}

// AbandonCh provides a mock function with given fields:
func (_m *MockStateStore) AbandonCh() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// NewMockStateStore creates a new instance of MockStateStore. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockStateStore(t testing.TB) *MockStateStore {
	mock := &MockStateStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}