// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	threadlikes "fgd-alterra-29/business/thread_likes"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetLikeState provides a mock function with given fields: ctx, domain, id
func (_m *Repository) GetLikeState(ctx context.Context, domain threadlikes.Domain, id int) (threadlikes.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 threadlikes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, threadlikes.Domain, int) threadlikes.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(threadlikes.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, threadlikes.Domain, int) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Like provides a mock function with given fields: ctx, domain, id
func (_m *Repository) Like(ctx context.Context, domain threadlikes.Domain, id int) (threadlikes.Domain, int, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 threadlikes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, threadlikes.Domain, int) threadlikes.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(threadlikes.Domain)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, threadlikes.Domain, int) int); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, threadlikes.Domain, int) error); ok {
		r2 = rf(ctx, domain, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewLike provides a mock function with given fields: ctx, domain, id
func (_m *Repository) NewLike(ctx context.Context, domain threadlikes.Domain, id int) (threadlikes.Domain, int, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 threadlikes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, threadlikes.Domain, int) threadlikes.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(threadlikes.Domain)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, threadlikes.Domain, int) int); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, threadlikes.Domain, int) error); ok {
		r2 = rf(ctx, domain, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Unlike provides a mock function with given fields: ctx, domain, id
func (_m *Repository) Unlike(ctx context.Context, domain threadlikes.Domain, id int) (threadlikes.Domain, int, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 threadlikes.Domain
	if rf, ok := ret.Get(0).(func(context.Context, threadlikes.Domain, int) threadlikes.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(threadlikes.Domain)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, threadlikes.Domain, int) int); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, threadlikes.Domain, int) error); ok {
		r2 = rf(ctx, domain, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
