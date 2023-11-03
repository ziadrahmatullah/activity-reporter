// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/entity"
	mock "github.com/stretchr/testify/mock"
)

// Notification is an autogenerated mock type for the Notification type
type Notification struct {
	mock.Mock
}

// Notify provides a mock function with given fields: _a0, _a1
func (_m *Notification) Notify(_a0 entity.Publisher, _a1 entity.Follower) string {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(entity.Publisher, entity.Follower) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewNotification interface {
	mock.TestingT
	Cleanup(func())
}

// NewNotification creates a new instance of Notification. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNotification(t mockConstructorTestingTNewNotification) *Notification {
	mock := &Notification{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}