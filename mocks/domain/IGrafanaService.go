// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "grafana-api/domain/model"
)

// IGrafanaService is an autogenerated mock type for the IGrafanaService type
type IGrafanaService struct {
	mock.Mock
}

// CheckOrganizationUser provides a mock function with given fields: ctx, cmd
func (_m *IGrafanaService) CheckOrganizationUser(ctx context.Context, cmd model.CheckOrgUserRequest) error {
	ret := _m.Called(ctx, cmd)

	if len(ret) == 0 {
		panic("no return value specified for CheckOrganizationUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CheckOrgUserRequest) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIGrafanaService creates a new instance of IGrafanaService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIGrafanaService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IGrafanaService {
	mock := &IGrafanaService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
