package userpoints_test

import (
	"context"
	"errors"
	up "fgd-alterra-29/business/user_points"
	_userpointMocks "fgd-alterra-29/business/user_points/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userpointRepository _userpointMocks.Repository
var userpointService up.UseCase
var userpointDomain up.Domain

func setup() {
	userpointService = up.NewUserPointUseCase(&userpointRepository, time.Hour*1)
	userpointDomain = up.Domain{
		User_id:     1,
		ThreadPoint: 50,
		PostPoint:   50,
	}
}

func TestAddThreadPointController(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		userpointRepository.On("AddThreadPoint", mock.Anything, mock.AnythingOfType("int")).Return(userpointDomain, nil).Once()

		userpoint, err := userpointService.AddThreadPointController(context.Background(), userpointDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, userpoint.User_id)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		userpointRepository.On("AddThreadPoint", mock.Anything, mock.AnythingOfType("int")).Return(userpointDomain, errors.New("")).Once()

		_, err := userpointService.AddThreadPointController(context.Background(), userpointDomain.User_id)
		assert.NotNil(t, err)
	})
}
func TestAddPostPointController(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		userpointRepository.On("AddPostPoint", mock.Anything, mock.AnythingOfType("int")).Return(userpointDomain, nil).Once()

		userpoint, err := userpointService.AddPostPointController(context.Background(), userpointDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, userpoint.User_id)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		userpointRepository.On("AddPostPoint", mock.Anything, mock.AnythingOfType("int")).Return(userpointDomain, errors.New("")).Once()

		_, err := userpointService.AddPostPointController(context.Background(), userpointDomain.User_id)
		assert.NotNil(t, err)
	})
}
