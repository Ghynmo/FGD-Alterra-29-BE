package threads_test

import (
	"context"
	"errors"
	"fgd-alterra-29/business/badges"
	"fgd-alterra-29/business/threads"
	_threadMocks "fgd-alterra-29/business/threads/mocks"
	userbadges "fgd-alterra-29/business/user_badges"

	// _userpointMocks "fgd-alterra-29/business/user_points/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var threadRepository _threadMocks.Repository

// var userpointRepository _userpointMocks.Repository
var threadService threads.UseCase
var threadDomain threads.Domain
var ubRepository userbadges.Repository
var badgeRepository badges.Repository

func setup() {
	threadService = threads.NewThreadUseCase(&threadRepository, time.Hour*1, badgeRepository, ubRepository)
	threadDomain = threads.Domain{
		Title:    "test",
		Q_Thread: 5,
	}
}

func TestGetThreadsByTitleController(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetThreadsByTitle", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadDomain}, nil).Once()

		thread, err := threadService.GetThreadsByTitleController(context.Background(), threadDomain.Title)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(thread))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetThreadsByTitle", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadDomain}, errors.New("")).Once()

		_, err := threadService.GetThreadsByTitleController(context.Background(), threadDomain.Title)
		assert.NotNil(t, err)
	})
}
func TestGetProfileThreads(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetProfileThreads", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadDomain}, nil).Once()

		thread, err := threadService.GetProfileThreads(context.Background(), threadDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(thread))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetProfileThreads", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadDomain}, errors.New("")).Once()

		_, err := threadService.GetProfileThreads(context.Background(), threadDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestGetThreadQuantity(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetThreadQuantity", mock.Anything).Return(threadDomain, nil).Once()

		thread, err := threadService.GetThreadQuantity(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 5, thread.Q_Thread)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetThreadQuantity", mock.Anything).Return(threadDomain, errors.New("")).Once()

		_, err := threadService.GetThreadQuantity(context.Background())
		assert.NotNil(t, err)
	})
}

// func TestCreateThread(t *testing.T) {
// 	setup()
// 	t.Run("Test Case 1 | Valid", func(t *testing.T) {
// 		threadRepository.On("CreateThread", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadDomain, nil).Once()
// 		userpointRepository.On("AddThreadPoint", mock.AnythingOfType("int")).Return(upDomain, nil)

// 		thread, err := threadService.CreateThread(context.Background(), threadDomain, threadDomain.ID)
// 		assert.Nil(t, err)
// 		assert.Equal(t, "test", thread.Title)
// 	})
// 	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
// 		threadRepository.On("CreateThread", mock.Anything, mock.AnythingOfType("threads.Domain"), mock.AnythingOfType("int")).Return(threadDomain, nil).Once()
// 		userpointRepository.On("AddThreadPoint", mock.AnythingOfType("int")).Return(upDomain, errors.New(""))

// 		thread, err := threadService.CreateThread(context.Background(), threadDomain, threadDomain.ID)
// 		assert.Nil(t, err)
// 		assert.Equal(t, "test", thread.Title)
// 	})
// 	t.Run("Test Case 3 | Invalid", func(t *testing.T) {
// 		threadRepository.On("CreateThread", mock.Anything, mock.AnythingOfType("threads.Domain"), mock.AnythingOfType("int")).Return(threadDomain, errors.New("")).Once()

// 		_, err := threadService.CreateThread(context.Background(), threadDomain, threadDomain.ID)
// 		assert.NotNil(t, err)
// 	})
// }
func TestGetHomepageThreads(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetHomepageThreads", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadDomain}, nil).Once()

		thread, err := threadService.GetHomepageThreads(context.Background(), threadDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(thread))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetHomepageThreads", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadDomain}, errors.New("")).Once()

		_, err := threadService.GetHomepageThreads(context.Background(), threadDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestGetThreads(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetThreads", mock.Anything).Return([]threads.Domain{threadDomain}, nil).Once()

		thread, err := threadService.GetThreads(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(thread))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetThreads", mock.Anything).Return([]threads.Domain{threadDomain}, errors.New("")).Once()

		_, err := threadService.GetThreads(context.Background())
		assert.NotNil(t, err)
	})
}
func TestGetRecommendationThreads(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetRecommendationThreads", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadDomain}, nil).Once()

		thread, err := threadService.GetRecommendationThreads(context.Background(), threadDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(thread))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetRecommendationThreads", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadDomain}, errors.New("")).Once()

		_, err := threadService.GetRecommendationThreads(context.Background(), threadDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestGetThreadByID(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetThreadByID", mock.Anything, mock.AnythingOfType("int")).Return(threadDomain, nil).Once()

		thread, err := threadService.GetThreadByID(context.Background(), threadDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, "test", thread.Title)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetThreadByID", mock.Anything, mock.AnythingOfType("int")).Return(threadDomain, errors.New("")).Once()

		_, err := threadService.GetThreadByID(context.Background(), threadDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestDeleteThread(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("DeleteThread", mock.Anything, mock.AnythingOfType("int")).Return(threadDomain, nil).Once()

		thread, err := threadService.DeleteThread(context.Background(), threadDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, "test", thread.Title)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("DeleteThread", mock.Anything, mock.AnythingOfType("int")).Return(threadDomain, errors.New("")).Once()

		_, err := threadService.DeleteThread(context.Background(), threadDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestActivateThread(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("ActivateThread", mock.Anything, mock.AnythingOfType("int")).Return(threadDomain, nil).Once()

		thread, err := threadService.ActivateThread(context.Background(), threadDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, "test", thread.Title)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("ActivateThread", mock.Anything, mock.AnythingOfType("int")).Return(threadDomain, errors.New("")).Once()

		_, err := threadService.ActivateThread(context.Background(), threadDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestGetHotThreads(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetHotThreads", mock.Anything).Return([]threads.Domain{threadDomain}, nil).Once()

		thread, err := threadService.GetHotThreads(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(thread))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetHotThreads", mock.Anything).Return([]threads.Domain{threadDomain}, errors.New("")).Once()

		_, err := threadService.GetHotThreads(context.Background())
		assert.NotNil(t, err)
	})
}
func TestGetSearch(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetSearch", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadDomain}, nil).Once()

		thread, err := threadService.GetSearch(context.Background(), threadDomain.Title)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(thread))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetSearch", mock.Anything, mock.AnythingOfType("string")).Return([]threads.Domain{threadDomain}, errors.New("")).Once()

		_, err := threadService.GetSearch(context.Background(), threadDomain.Title)
		assert.NotNil(t, err)
	})
}
func TestGetThreadsByCategoryID(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetThreadsByCategoryID", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadDomain}, nil).Once()

		thread, err := threadService.GetThreadsByCategoryID(context.Background(), threadDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(thread))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetThreadsByCategoryID", mock.Anything, mock.AnythingOfType("int")).Return([]threads.Domain{threadDomain}, errors.New("")).Once()

		_, err := threadService.GetThreadsByCategoryID(context.Background(), threadDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestGetSideNewsThreads(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadRepository.On("GetSideNewsThreads", mock.Anything).Return([]threads.Domain{threadDomain}, nil).Once()

		thread, err := threadService.GetSideNewsThreads(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(thread))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadRepository.On("GetSideNewsThreads", mock.Anything).Return([]threads.Domain{threadDomain}, errors.New("")).Once()

		_, err := threadService.GetSideNewsThreads(context.Background())
		assert.NotNil(t, err)
	})
}
