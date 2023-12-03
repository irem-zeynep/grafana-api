// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "grafana-api/domain/model"
)

// IAuditService is an autogenerated mock type for the IAuditService type
type IAuditService struct {
	mock.Mock
}

// SaveAudit provides a mock function with given fields: ctx, dto
func (_m *IAuditService) SaveAudit(ctx context.Context, dto model.AuditDTO) error {
	ret := _m.Called(ctx, dto)

	if len(ret) == 0 {
		panic("no return value specified for SaveAudit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.AuditDTO) error); ok {
		r0 = rf(ctx, dto)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIAuditService creates a new instance of IAuditService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuditService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuditService {
	mock := &IAuditService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
