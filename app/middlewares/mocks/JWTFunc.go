// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt/v4"
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// JWTFunc is an autogenerated mock type for the JWTFunc type
type JWTFunc struct {
	mock.Mock
}

// ExtractAdmin provides a mock function with given fields: c
func (_m *JWTFunc) ExtractAdmin(c echo.Context) {
	_m.Called(c)
}

// ExtractClaims provides a mock function with given fields: tokenStr
func (_m *JWTFunc) ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	ret := _m.Called(tokenStr)

	var r0 jwt.MapClaims
	if rf, ok := ret.Get(0).(func(string) jwt.MapClaims); ok {
		r0 = rf(tokenStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jwt.MapClaims)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(tokenStr)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// ExtractID provides a mock function with given fields: c
func (_m *JWTFunc) ExtractID(c echo.Context) int {
	ret := _m.Called(c)

	var r0 int
	if rf, ok := ret.Get(0).(func(echo.Context) int); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GenerateToken provides a mock function with given fields: id, admin
func (_m *JWTFunc) GenerateToken(id int, admin bool) (string, error) {
	ret := _m.Called(id, admin)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, bool) string); ok {
		r0 = rf(id, admin)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, bool) error); ok {
		r1 = rf(id, admin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
