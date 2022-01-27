package reputations_test

import (
	"context"
	"errors"
	"fgd-alterra-29/business/reputations"
	_reputationMocks "fgd-alterra-29/business/reputations/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var reputationRepository _reputationMocks.Repository
var reputationService reputations.UseCase
var reputationDomain reputations.Domain

func setup() {
	reputationService = reputations.NewReputationUseCase(&reputationRepository, time.Hour*1)
	reputationDomain = reputations.Domain{
		ID: 1,
	}
}

func TestCreateReputation(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		reputationRepository.On("CreateReputation", mock.Anything, mock.Anything).Return(reputationDomain, nil).Once()

		reputation, err := reputationService.CreateReputationController(context.Background(), reputationDomain)
		assert.Nil(t, err)
		assert.Equal(t, 1, reputation.ID)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		reputationRepository.On("CreateReputation", mock.Anything, mock.Anything).Return(reputationDomain, errors.New("")).Once()

		_, err := reputationService.CreateReputationController(context.Background(), reputationDomain)
		assert.NotNil(t, err)
	})
}

func TestGetReputationByUser(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		reputationRepository.On("GetReputationByUser", mock.Anything, mock.AnythingOfType("int")).Return(reputationDomain, nil).Once()

		reputation, err := reputationService.GetReputationByUser(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, reputation.ID)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		reputationRepository.On("GetReputationByUser", mock.Anything, mock.AnythingOfType("int")).Return(reputationDomain, errors.New("")).Once()

		_, err := reputationService.GetReputationByUser(context.Background(), 1)
		assert.NotNil(t, err)
	})
}
