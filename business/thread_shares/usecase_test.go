package threadshares_test

import (
	"context"
	"errors"
	ts "fgd-alterra-29/business/thread_shares"
	_threadshareMocks "fgd-alterra-29/business/thread_shares/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var threadshareRepository _threadshareMocks.Repository
var threadshareService ts.UseCase
var threadshareDomain ts.Domain

func setup() {
	threadshareService = ts.NewThreadShareUseCase(&threadshareRepository, time.Hour*1)
	threadshareDomain = ts.Domain{
		User_id:   1,
		Shared_at: time.Time{},
	}
}

func TestGetThreadShare(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadshareRepository.On("ThreadShare", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadshareDomain, nil).Once()

		threadshare, err := threadshareService.ThreadShareController(context.Background(), threadshareDomain, threadshareDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, threadshare.User_id)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadshareRepository.On("ThreadShare", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadshareDomain, errors.New("")).Once()

		_, err := threadshareService.ThreadShareController(context.Background(), threadshareDomain, threadshareDomain.User_id)
		assert.NotNil(t, err)
	})
}
