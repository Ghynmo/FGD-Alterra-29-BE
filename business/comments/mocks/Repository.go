// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	comments "fgd-alterra-29/business/comments"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ActivatingPost provides a mock function with given fields: ctx, id
func (_m *Repository) ActivatingPost(ctx context.Context, id int) (comments.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) comments.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(comments.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateComment provides a mock function with given fields: ctx, domain, id
func (_m *Repository) CreateComment(ctx context.Context, domain comments.Domain, id int) (comments.Domain, error) {
	ret := _m.Called(ctx, domain, id)

	var r0 comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, comments.Domain, int) comments.Domain); ok {
		r0 = rf(ctx, domain, id)
	} else {
		r0 = ret.Get(0).(comments.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, comments.Domain, int) error); ok {
		r1 = rf(ctx, domain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommentByThread provides a mock function with given fields: ctx, thread_id, my_id
func (_m *Repository) GetCommentByThread(ctx context.Context, thread_id int, my_id int) ([]comments.Domain, error) {
	ret := _m.Called(ctx, thread_id, my_id)

	var r0 []comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []comments.Domain); ok {
		r0 = rf(ctx, thread_id, my_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, thread_id, my_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommentProfile provides a mock function with given fields: ctx, id
func (_m *Repository) GetCommentProfile(ctx context.Context, id int) ([]comments.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 []comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) []comments.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Domain)
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

// GetCommentReply provides a mock function with given fields: ctx, id, reply_of
func (_m *Repository) GetCommentReply(ctx context.Context, id int, reply_of int) ([]comments.Domain, error) {
	ret := _m.Called(ctx, id, reply_of)

	var r0 []comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []comments.Domain); ok {
		r0 = rf(ctx, id, reply_of)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, id, reply_of)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostQuantity provides a mock function with given fields: ctx
func (_m *Repository) GetPostQuantity(ctx context.Context) (comments.Domain, error) {
	ret := _m.Called(ctx)

	var r0 comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context) comments.Domain); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(comments.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPosts provides a mock function with given fields: ctx
func (_m *Repository) GetPosts(ctx context.Context) ([]comments.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []comments.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostsByComment provides a mock function with given fields: ctx, comment
func (_m *Repository) GetPostsByComment(ctx context.Context, comment string) ([]comments.Domain, error) {
	ret := _m.Called(ctx, comment)

	var r0 []comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) []comments.Domain); ok {
		r0 = rf(ctx, comment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, comment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnactivatingPost provides a mock function with given fields: ctx, id
func (_m *Repository) UnactivatingPost(ctx context.Context, id int) (comments.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 comments.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) comments.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(comments.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
