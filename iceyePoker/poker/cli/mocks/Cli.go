// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Cli is an autogenerated mock type for the Cli type
type Cli struct {
	mock.Mock
}

// AskBoolQuestion provides a mock function with given fields: message
func (_m *Cli) AskBoolQuestion(message string) (bool, error) {
	ret := _m.Called(message)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadHand provides a mock function with given fields: message
func (_m *Cli) ReadHand(message string) (string, error) {
	ret := _m.Called(message)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCli interface {
	mock.TestingT
	Cleanup(func())
}

// NewCli creates a new instance of Cli. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCli(t mockConstructorTestingTNewCli) *Cli {
	mock := &Cli{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
