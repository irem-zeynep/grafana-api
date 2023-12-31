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

// CheckOrganizationUser provides a mock function with given fields: ctx, req
func (_m *IGrafanaService) CheckOrganizationUser(ctx context.Context, req model.CheckOrgUserRequest) (*model.CheckOrgUserDTO, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CheckOrganizationUser")
	}

	var r0 *model.CheckOrgUserDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CheckOrgUserRequest) (*model.CheckOrgUserDTO, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CheckOrgUserRequest) *model.CheckOrgUserDTO); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CheckOrgUserDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CheckOrgUserRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
