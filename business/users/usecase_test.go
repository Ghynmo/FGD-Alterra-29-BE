package users_test

import (
	"context"
	"errors"
	"fgd-alterra-29/app/middlewares"
	"fgd-alterra-29/business/users"
	_userMocks "fgd-alterra-29/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _userMocks.Repository
var userService users.UseCase
var userDomain users.Domain
var jwtAuth middlewares.ConfigJWT

func setup() {
	userService = users.NewUserUseCase(&userRepository, time.Hour*1, jwtAuth)
	userDomain = users.Domain{
		Name:     "tester",
		Email:    "test@mail.com",
		Password: "PassHashed",
	}
}

func TestRegisterController(t *testing.T) {
	setup()
	t.Run("Test RegisterController | Valid but Avoid Hash", func(t *testing.T) {
		userRepository.On("CheckUsername",
			mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepository.On("CheckEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepository.On("Register",
			mock.Anything, mock.AnythingOfType("users.Domain")).Return(userDomain, errors.New("")).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
	})

	t.Run("Test RegisterController | Invalid (Username Has been used)", func(t *testing.T) {
		userRepository.On("CheckUsername",
			mock.Anything, mock.AnythingOfType("string")).Return(true, nil).Once()
		userRepository.On("CheckEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
		// userRepository.On("Register",
		// 	mock.Anything, mock.AnythingOfType("users.Domain")).Return(userDomain, errors.New("")).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
		assert.Equal(t, "USERNAME IS ALREADY USED", err.Error())
	})

	t.Run("Test RegisterController | Invalid (Username Has been used)", func(t *testing.T) {
		userRepository.On("CheckUsername",
			mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepository.On("CheckEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(true, nil).Once()
		// userRepository.On("Register",
		// 	mock.Anything, mock.AnythingOfType("users.Domain")).Return(userDomain, errors.New("")).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
		assert.Equal(t, "EMAIL IS ALREADY USED", err.Error())
	})

	t.Run("Test RegisterController | Invalid (Username Has been used)", func(t *testing.T) {
		userRepository.On("CheckUsername",
			mock.Anything, mock.AnythingOfType("string")).Return(true, nil).Once()
		userRepository.On("CheckEmail",
			mock.Anything, mock.AnythingOfType("string")).Return(true, nil).Once()
		// userRepository.On("Register",
		// 	mock.Anything, mock.AnythingOfType("users.Domain")).Return(userDomain, errors.New("")).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
		assert.Equal(t, "USERNAME & EMAIL IS ALREADY USED", err.Error())
	})

	t.Run("Test RegisterController | Invalid (Empty Name)", func(t *testing.T) {
		userDomain.Name = ""
		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.Equal(t, "NAME MUST BE FILLED", err.Error())
	})

	t.Run("Test RegisterController | Invalid (Empty Email)", func(t *testing.T) {
		userDomain.Name = "tester"
		userDomain.Email = ""
		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.Equal(t, "EMAIL MUST BE FILLED", err.Error())
	})

	t.Run("Test RegisterController | Invalid (Empty Password)", func(t *testing.T) {
		userDomain.Name = "tester"
		userDomain.Email = "test@mail.com"
		userDomain.Password = ""
		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.Equal(t, "PASSWORD MUST BE FILLED", err.Error())
	})
}

func TestGetUsersController(t *testing.T) {
	setup()
	t.Run("Test case 1 | GetUsersController", func(t *testing.T) {
		userRepository.On("GetUsers",
			mock.Anything,
		// mock.AnythingOfType("users.Domain")
		).
			Return([]users.Domain{userDomain}, nil).Once()

		// userRepository.On("CheckUsername",
		// 	mock.AnythingOfType("context.Context"),
		// 	mock.AnythingOfType("string")).
		// 	Return(userDomain, nil).Once()

		user, err := userService.GetUsersController(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(user))
	})
	t.Run("Test case 2 | GetUsersController", func(t *testing.T) {
		userRepository.On("GetUsers",
			mock.Anything,
		// mock.AnythingOfType("users.Domain")
		).
			Return([]users.Domain{}, errors.New("")).Once()

		_, err := userService.GetUsersController(context.Background())
		assert.NotNil(t, err)
	})
}
