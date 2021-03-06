// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	badges "fgd-alterra-29/business/badges"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ActivateBadge provides a mock function with given fields: ctx, domain
func (_m *Repository) ActivateBadge(ctx context.Context, domain badges.Domain) (badges.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 badges.Domain
	if rf, ok := ret.Get(0).(func(context.Context, badges.Domain) badges.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(badges.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, badges.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBadge provides a mock function with given fields: ctx, domain
func (_m *Repository) CreateBadge(ctx context.Context, domain badges.Domain) (badges.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 badges.Domain
	if rf, ok := ret.Get(0).(func(context.Context, badges.Domain) badges.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(badges.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, badges.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBadgesByUser provides a mock function with given fields: ctx, id
func (_m *Repository) GetBadgesByUser(ctx context.Context, id int) ([]badges.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 []badges.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) []badges.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]badges.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBadgesIdByThread provides a mock function with given fields: ctx, thread_qty
func (_m *Repository) GetBadgesIdByThread(ctx context.Context, thread_qty int) (int, error) {
	ret := _m.Called(ctx, thread_qty)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, int) int); ok {
		r0 = rf(ctx, thread_qty)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, thread_qty)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnactivateBadge provides a mock function with given fields: ctx, domain
func (_m *Repository) UnactivateBadge(ctx context.Context, domain badges.Domain) (badges.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 badges.Domain
	if rf, ok := ret.Get(0).(func(context.Context, badges.Domain) badges.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(badges.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, badges.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
