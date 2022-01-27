package badges_test

import (
	"context"
	"errors"
	"fgd-alterra-29/business/badges"
	_badgeMocks "fgd-alterra-29/business/badges/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var badgeRepository _badgeMocks.Repository
var badgeService badges.UseCase
var badgeDomain badges.Domain

func setup() {
	badgeService = badges.NewBadgeUseCase(&badgeRepository, time.Hour*1)
	badgeDomain = badges.Domain{
		ID:     1,
		Badge:  "Social",
		Status: true,
	}
}

func TestGetBadgesByUser(t *testing.T) {
	setup()
	UserID := 1
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		badgeRepository.On("GetBadgesByUser", mock.Anything, mock.AnythingOfType("int")).Return([]badges.Domain{badgeDomain}, nil).Once()

		badge, err := badgeService.GetBadgesByUserController(context.Background(), UserID)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(badge))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		badgeRepository.On("GetBadgesByUser", mock.Anything, mock.AnythingOfType("int")).Return([]badges.Domain{badgeDomain}, errors.New("")).Once()

		_, err := badgeService.GetBadgesByUserController(context.Background(), UserID)
		assert.NotNil(t, err)
	})
}

func TestCreateBadge(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		badgeRepository.On("CreateBadge", mock.Anything, mock.Anything).Return(badgeDomain, nil).Once()

		badge, err := badgeService.CreateBadgeController(context.Background(), badgeDomain)
		assert.Nil(t, err)
		assert.Equal(t, 1, badge.ID)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		badgeRepository.On("CreateBadge", mock.Anything, mock.Anything).Return(badgeDomain, errors.New("")).Once()

		_, err := badgeService.CreateBadgeController(context.Background(), badgeDomain)
		assert.NotNil(t, err)
	})
}
