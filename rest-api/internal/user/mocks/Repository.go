// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"
	auth "rest-api/m/rest-api/internal/auth"

	mock "github.com/stretchr/testify/mock"

	user "rest-api/m/rest-api/internal/user"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *Repository) GetByEmail(ctx context.Context, email string) (auth.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 auth.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (auth.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) auth.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(auth.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, _a1
func (_m *Repository) GetById(ctx context.Context, _a1 *user.GetUser) (user.User, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUser) (user.User, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUser) user.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetUser) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *Repository) Update(ctx context.Context, _a1 *user.UpdateUser) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.UpdateUser) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
