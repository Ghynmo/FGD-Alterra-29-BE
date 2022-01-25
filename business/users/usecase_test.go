package users_test

import (
	"context"
	"errors"
	"fgd-alterra-29/app/middlewares"
	"fgd-alterra-29/business/users"
	_userMocks "fgd-alterra-29/business/users/mocks"
	"fgd-alterra-29/helpers"

	// _bycriptMocks "fgd-alterra-29/helpers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _userMocks.Repository
var userService users.UseCase
var userDomain users.Domain

// var helperBycript _bycriptMocks.Bycript

// var jwtAuth _jwtMocks.JWTFunc

func setup() {
	userService = users.NewUserUseCase(&userRepository, time.Hour*1, middlewares.ConfigJWT{})
	hashedPassword, _ := helpers.Hash("test")
	userDomain = users.Domain{
		ID:       1,
		Name:     "tester",
		Email:    "test@mail.com",
		Password: hashedPassword,
		Q_User:   1,
		Status:   "active",
		Token:    "GeneratedToken",
	}
}

func TestRegisterController(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid (Avoid Token)", func(t *testing.T) {
		userRepository.On("CheckUsername", mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepository.On("CheckEmail", mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepository.On("Register", mock.Anything, mock.AnythingOfType("users.Domain")).Return(userDomain, errors.New("")).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.Equal(t, "", err.Error())
	})
	t.Run("Test Case 2 | Invalid (Avoid Token)", func(t *testing.T) {
		userRepository.On("CheckUsername", mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepository.On("CheckEmail", mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepository.On("Register", mock.Anything, mock.AnythingOfType("users.Domain")).Return(userDomain, errors.New("")).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid (Error CheckName)", func(t *testing.T) {
		userRepository.On("CheckUsername", mock.Anything, mock.AnythingOfType("string")).Return(true, errors.New("")).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Invalid (Error CheckEmail)", func(t *testing.T) {
		userRepository.On("CheckUsername", mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepository.On("CheckEmail", mock.Anything, mock.AnythingOfType("string")).Return(true, errors.New("")).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
	})

	t.Run("Test Case 5 | Invalid (Username Has been used)", func(t *testing.T) {
		userRepository.On("CheckUsername", mock.Anything, mock.AnythingOfType("string")).Return(true, nil).Once()
		userRepository.On("CheckEmail", mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
		assert.Equal(t, "USERNAME IS ALREADY USED", err.Error())
	})

	t.Run("Test Case 6 | Invalid (Email Has been used)", func(t *testing.T) {
		userRepository.On("CheckUsername", mock.Anything, mock.AnythingOfType("string")).Return(false, nil).Once()
		userRepository.On("CheckEmail", mock.Anything, mock.AnythingOfType("string")).Return(true, nil).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
		assert.Equal(t, "EMAIL IS ALREADY USED", err.Error())
	})

	t.Run("Test Case 7 | Invalid (Username & Email Has been used)", func(t *testing.T) {
		userRepository.On("CheckUsername", mock.Anything, mock.AnythingOfType("string")).Return(true, nil).Once()
		userRepository.On("CheckEmail", mock.Anything, mock.AnythingOfType("string")).Return(true, nil).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
		assert.Equal(t, "USERNAME & EMAIL IS ALREADY USED", err.Error())
	})

	t.Run("Test Case 8 | Invalid (Empty Name)", func(t *testing.T) {
		userDomain.Name = ""

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.Equal(t, "NAME MUST BE FILLED", err.Error())
	})

	t.Run("Test Case 9 | Invalid (Empty Email)", func(t *testing.T) {
		userDomain.Name = "tester"
		userDomain.Email = ""

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.Equal(t, "EMAIL MUST BE FILLED", err.Error())
	})

	t.Run("Test Case 10 | Invalid (Empty Password)", func(t *testing.T) {
		userDomain.Name = "tester"
		userDomain.Email = "test@mail.com"
		userDomain.Password = ""

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.Equal(t, "PASSWORD MUST BE FILLED", err.Error())
	})
	t.Run("Test RegisterController | Valid but Admin", func(t *testing.T) {
		userRepository.On("Register", mock.Anything, mock.Anything).Return(users.Domain{Role_id: 1}, nil).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
	})
	t.Run("Test RegisterController | Valid but non Admin", func(t *testing.T) {
		userRepository.On("Register", mock.Anything, mock.Anything).Return(users.Domain{Role_id: 2}, nil).Once()

		_, err := userService.RegisterController(context.Background(), userDomain)
		assert.NotNil(t, err)
	})
}

func TestLoginController(t *testing.T) {
	setup()
	t.Run("Test LoginController | Valid but avoid token)", func(t *testing.T) {
		userDomain.Email = "test@mail.com"
		userDomain.Password = "HashedPassword"
		userRepository.On("Login", mock.Anything, mock.AnythingOfType("users.Domain")).Return(userDomain, errors.New("")).Once()

		_, err := userService.LoginController(context.Background(), userDomain)
		assert.NotNil(t, err)
	})

	t.Run("Test LoginController | Invalid (Empty Email)", func(t *testing.T) {
		userDomain.Email = ""

		_, err := userService.LoginController(context.Background(), userDomain)
		assert.Equal(t, "EMAIL MUST BE FILLED", err.Error())
	})

	t.Run("Test LoginController | Invalid (Empty Password)", func(t *testing.T) {
		userDomain.Email = "test@mail.com"
		userDomain.Password = ""

		_, err := userService.LoginController(context.Background(), userDomain)
		assert.Equal(t, "PASSWORD MUST BE FILLED", err.Error())
	})
	t.Run("Test LoginController | Valid but avoid token", func(t *testing.T) {
		userDomain.Email = "test@mail.com"
		userDomain.Password = "HashedPassword"
		userRepository.On("Login", mock.Anything, mock.AnythingOfType("users.Domain")).Return(users.Domain{Role_id: 1}, nil).Once()

		_, err := userService.LoginController(context.Background(), userDomain)
		assert.NotNil(t, err)
	})
	t.Run("Test LoginController | Valid but avoid token", func(t *testing.T) {
		userDomain.Email = "test@mail.com"
		userDomain.Password = "HashedPassword"
		userRepository.On("Login", mock.Anything, mock.AnythingOfType("users.Domain")).Return(users.Domain{Role_id: 2}, nil).Once()

		_, err := userService.LoginController(context.Background(), userDomain)
		assert.NotNil(t, err)
	})
}

func TestGetUsersController(t *testing.T) {
	setup()
	t.Run("Test case 1 | GetUsersController", func(t *testing.T) {
		userRepository.On("GetUsers", mock.Anything).Return([]users.Domain{userDomain}, nil).Once()

		user, err := userService.GetUsersController(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(user))
	})
	t.Run("Test case 2 | GetUsersController", func(t *testing.T) {
		userRepository.On("GetUsers", mock.Anything).Return([]users.Domain{}, errors.New("")).Once()

		_, err := userService.GetUsersController(context.Background())
		assert.NotNil(t, err)
	})
}

func TestGetUsersByNameController(t *testing.T) {
	setup()
	t.Run("Test case 1 | GetUsersByNameController", func(t *testing.T) {
		userRepository.On("GetUsersByName", mock.Anything, mock.AnythingOfType("string")).Return([]users.Domain{userDomain}, nil).Once()

		user, err := userService.GetUsersByNameController(context.Background(), userDomain.Name)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(user))
	})
	t.Run("Test case 2 | GetUsersByNameController", func(t *testing.T) {
		userRepository.On("GetUsersByName", mock.Anything, mock.AnythingOfType("string")).Return([]users.Domain{}, errors.New("")).Once()

		_, err := userService.GetUsersByNameController(context.Background(), userDomain.Name)
		assert.NotNil(t, err)
	})
}

func TestGetProfileController(t *testing.T) {
	setup()
	t.Run("Test case 1 | GetProfileController", func(t *testing.T) {
		userRepository.On("GetProfile", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()

		user, err := userService.GetProfileController(context.Background(), userDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, user.ID)
	})
	t.Run("Test case 2 | GetProfileController", func(t *testing.T) {
		userRepository.On("GetProfile", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("")).Once()

		_, err := userService.GetProfileController(context.Background(), userDomain.ID)
		assert.NotNil(t, err)
	})
}

func TestGetUsersQuantity(t *testing.T) {
	setup()
	t.Run("Test case 1 | GetUsersQuantity", func(t *testing.T) {
		userRepository.On("GetUsersQuantity", mock.Anything).Return(userDomain, nil).Once()

		user, err := userService.GetUsersQuantity(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, user.Q_User)
	})
	t.Run("Test case 2 | GetUsersQuantity", func(t *testing.T) {
		userRepository.On("GetUsersQuantity", mock.Anything).Return(users.Domain{}, errors.New("")).Once()

		_, err := userService.GetUsersQuantity(context.Background())
		assert.NotNil(t, err)
	})
}

func TestGetProfileSetting(t *testing.T) {
	setup()
	t.Run("Test case 1 | GetProfileSetting", func(t *testing.T) {
		userRepository.On("GetProfileSetting", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()

		user, err := userService.GetProfileSetting(context.Background(), userDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, "tester", user.Name)
	})
	t.Run("Test case 2 | GetProfileSetting", func(t *testing.T) {
		userRepository.On("GetProfileSetting", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("")).Once()

		_, err := userService.GetProfileSetting(context.Background(), userDomain.ID)
		assert.NotNil(t, err)
	})
}

func TestUpdateProfile(t *testing.T) {
	setup()
	t.Run("Test case 1 | UpdateProfile", func(t *testing.T) {
		userRepository.On("UpdateProfile", mock.Anything, mock.AnythingOfType("users.Domain"), mock.AnythingOfType("int")).Return(userDomain, nil).Once()

		user, err := userService.UpdateProfile(context.Background(), userDomain, userDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, "tester", user.Name)
	})
	t.Run("Test case 2 | UpdateProfile", func(t *testing.T) {
		userRepository.On("UpdateProfile", mock.Anything, mock.AnythingOfType("users.Domain"), mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("")).Once()

		_, err := userService.UpdateProfile(context.Background(), userDomain, userDomain.ID)
		assert.NotNil(t, err)
	})
}

func TestBannedUserController(t *testing.T) {
	setup()
	t.Run("Test case 1 | BannedUserController", func(t *testing.T) {
		userRepository.On("GetBannedState", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{Status: "active"}, nil).Once()
		userRepository.On("BannedUser", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()

		_, err := userService.BannedUserController(context.Background(), userDomain.ID)
		assert.Nil(t, err)
	})
	t.Run("Test case 2 | BannedUserController", func(t *testing.T) {
		userRepository.On("GetBannedState", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepository.On("BannedUser", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{Status: "active"}, errors.New("Error")).Once()

		_, err := userService.BannedUserController(context.Background(), userDomain.ID)
		assert.NotNil(t, err)
	})
	t.Run("Test case 3 | BannedUserController", func(t *testing.T) {
		userRepository.On("GetBannedState", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{Status: "banned"}, nil).Once()
		userRepository.On("UnbannedUser", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()

		_, err := userService.BannedUserController(context.Background(), userDomain.ID)
		assert.Nil(t, err)
	})
	t.Run("Test case 4 | BannedUserController", func(t *testing.T) {
		userRepository.On("GetBannedState", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{Status: "banned"}, nil).Once()
		userRepository.On("UnbannedUser", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("Error")).Once()

		_, err := userService.BannedUserController(context.Background(), userDomain.ID)
		assert.NotNil(t, err)
	})
	t.Run("Test case 5 | BannedUserController", func(t *testing.T) {
		userRepository.On("GetBannedState", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{Status: "not active or banned"}, nil).Once()

		user, err := userService.BannedUserController(context.Background(), userDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, users.Domain{}, user)
	})
}
