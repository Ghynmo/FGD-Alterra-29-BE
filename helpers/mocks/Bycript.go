// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Bycript is an autogenerated mock type for the Bycript type
type Bycript struct {
	mock.Mock
}

// Hash provides a mock function with given fields: password
func (_m *Bycript) Hash(password string) (string, error) {
	ret := _m.Called(password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateHash provides a mock function with given fields: password, hash
func (_m *Bycript) ValidateHash(password string, hash string) bool {
	ret := _m.Called(password, hash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(password, hash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
