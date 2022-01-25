package userbadges_test

import (
	"context"
	"errors"
	ub "fgd-alterra-29/business/user_badges"
	_userbadgeMocks "fgd-alterra-29/business/user_badges/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userbadgeRepository _userbadgeMocks.Repository
var userbadgeService ub.UseCase
var userbadgeDomain ub.Domain

func setup() {
	userbadgeService = ub.NewUserBadgeUseCase(&userbadgeRepository, time.Hour*1)
	userbadgeDomain = ub.Domain{
		User_id: 1,
	}
}

func TestGetUserBadge(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		userbadgeRepository.On("GetUserBadge", mock.Anything, mock.AnythingOfType("int")).Return([]ub.Domain{userbadgeDomain}, nil).Once()

		userbadge, err := userbadgeService.GetUserBadge(context.Background(), userbadgeDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(userbadge))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		userbadgeRepository.On("GetUserBadge", mock.Anything, mock.AnythingOfType("int")).Return([]ub.Domain{userbadgeDomain}, errors.New("")).Once()

		_, err := userbadgeService.GetUserBadge(context.Background(), userbadgeDomain.User_id)
		assert.NotNil(t, err)
	})
}
