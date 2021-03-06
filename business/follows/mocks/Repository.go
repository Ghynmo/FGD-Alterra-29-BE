// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	follows "fgd-alterra-29/business/follows"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Follows provides a mock function with given fields: ctx, domain, my_id
func (_m *Repository) Follows(ctx context.Context, domain follows.Domain, my_id int) (follows.Domain, error) {
	ret := _m.Called(ctx, domain, my_id)

	var r0 follows.Domain
	if rf, ok := ret.Get(0).(func(context.Context, follows.Domain, int) follows.Domain); ok {
		r0 = rf(ctx, domain, my_id)
	} else {
		r0 = ret.Get(0).(follows.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, follows.Domain, int) error); ok {
		r1 = rf(ctx, domain, my_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFollowState provides a mock function with given fields: ctx, domain, my_id
func (_m *Repository) GetFollowState(ctx context.Context, domain follows.Domain, my_id int) (follows.Domain, error) {
	ret := _m.Called(ctx, domain, my_id)

	var r0 follows.Domain
	if rf, ok := ret.Get(0).(func(context.Context, follows.Domain, int) follows.Domain); ok {
		r0 = rf(ctx, domain, my_id)
	} else {
		r0 = ret.Get(0).(follows.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, follows.Domain, int) error); ok {
		r1 = rf(ctx, domain, my_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFollowers provides a mock function with given fields: ctx, target_id, my_id
func (_m *Repository) GetFollowers(ctx context.Context, target_id int, my_id int) ([]follows.Domain, error) {
	ret := _m.Called(ctx, target_id, my_id)

	var r0 []follows.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []follows.Domain); ok {
		r0 = rf(ctx, target_id, my_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]follows.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, target_id, my_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFollowing provides a mock function with given fields: ctx, target_id, my_id
func (_m *Repository) GetFollowing(ctx context.Context, target_id int, my_id int) ([]follows.Domain, error) {
	ret := _m.Called(ctx, target_id, my_id)

	var r0 []follows.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []follows.Domain); ok {
		r0 = rf(ctx, target_id, my_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]follows.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, target_id, my_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFollow provides a mock function with given fields: ctx, domain, my_id
func (_m *Repository) NewFollow(ctx context.Context, domain follows.Domain, my_id int) (follows.Domain, error) {
	ret := _m.Called(ctx, domain, my_id)

	var r0 follows.Domain
	if rf, ok := ret.Get(0).(func(context.Context, follows.Domain, int) follows.Domain); ok {
		r0 = rf(ctx, domain, my_id)
	} else {
		r0 = ret.Get(0).(follows.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, follows.Domain, int) error); ok {
		r1 = rf(ctx, domain, my_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unfollow provides a mock function with given fields: ctx, domain, my_id
func (_m *Repository) Unfollow(ctx context.Context, domain follows.Domain, my_id int) (follows.Domain, error) {
	ret := _m.Called(ctx, domain, my_id)

	var r0 follows.Domain
	if rf, ok := ret.Get(0).(func(context.Context, follows.Domain, int) follows.Domain); ok {
		r0 = rf(ctx, domain, my_id)
	} else {
		r0 = ret.Get(0).(follows.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, follows.Domain, int) error); ok {
		r1 = rf(ctx, domain, my_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
