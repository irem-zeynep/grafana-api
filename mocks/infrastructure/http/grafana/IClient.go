// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"
	grafana "grafana-api/infrastructure/http/grafana"

	mock "github.com/stretchr/testify/mock"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, cmd
func (_m *IClient) CreateUser(ctx context.Context, cmd grafana.CreateUserCommand) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, grafana.CreateUserCommand) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOrg provides a mock function with given fields: ctx, name
func (_m *IClient) GetOrg(ctx context.Context, name string) (*grafana.OrganizationDTO, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetOrg")
	}

	var r0 *grafana.OrganizationDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*grafana.OrganizationDTO, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *grafana.OrganizationDTO); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grafana.OrganizationDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, name
func (_m *IClient) GetUser(ctx context.Context, name string) (*grafana.UserDTO, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 *grafana.UserDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*grafana.UserDTO, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *grafana.UserDTO); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grafana.UserDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIClient creates a new instance of IClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *IClient {
	mock := &IClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}