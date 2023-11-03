// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-activity-reporter/entity"
	mock "github.com/stretchr/testify/mock"
)

// Publisher is an autogenerated mock type for the Publisher type
type Publisher struct {
	mock.Mock
}

// FollowedBy provides a mock function with given fields: _a0
func (_m *Publisher) FollowedBy(_a0 entity.Follower) {
	_m.Called(_a0)
}

// PublisherNotificationAboutLike provides a mock function with given fields: _a0, _a1
func (_m *Publisher) PublisherNotificationAboutLike(_a0 entity.Follower, _a1 entity.Notification) {
	_m.Called(_a0, _a1)
}

// UploadPhoto provides a mock function with given fields:
func (_m *Publisher) UploadPhoto() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserName provides a mock function with given fields:
func (_m *Publisher) UserName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UserPhoto provides a mock function with given fields:
func (_m *Publisher) UserPhoto() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// addLiker provides a mock function with given fields: _a0
func (_m *Publisher) addLiker(_a0 entity.Follower) {
	_m.Called(_a0)
}

// notifyActivityToFollowers provides a mock function with given fields: _a0, _a1
func (_m *Publisher) notifyActivityToFollowers(_a0 string, _a1 entity.Publisher) {
	_m.Called(_a0, _a1)
}

// notifyUploadToFollowers provides a mock function with given fields: _a0
func (_m *Publisher) notifyUploadToFollowers(_a0 entity.Notification) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewPublisher interface {
	mock.TestingT
	Cleanup(func())
}

// NewPublisher creates a new instance of Publisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPublisher(t mockConstructorTestingTNewPublisher) *Publisher {
	mock := &Publisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}