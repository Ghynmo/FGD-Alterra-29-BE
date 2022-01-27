package threadlikes_test

import (
	"context"
	"errors"
	tl "fgd-alterra-29/business/thread_likes"
	_threadlikeMocks "fgd-alterra-29/business/thread_likes/mocks"
	userpoint "fgd-alterra-29/business/user_points"
	_userpointMocks "fgd-alterra-29/business/user_points/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var threadlikeRepository _threadlikeMocks.Repository
var threadlikeService tl.UseCase
var threadlikeDomain tl.Domain
var userpointDomain userpoint.Domain
var upMockRepository _userpointMocks.Repository

func setup() {
	threadlikeService = tl.NewThreadLikeUseCase(&threadlikeRepository, time.Hour*1, &upMockRepository)
	threadlikeDomain = tl.Domain{
		User_id:   1,
		Thread_id: 1,
		Liked_at:  time.Time{},
	}
}

func TestGetLikeState(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{User_id: 0}, nil).Once()
		threadlikeRepository.On("NewLike", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadlikeDomain, 1, nil).Once()
		upMockRepository.On("AddReputationPoint", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(userpointDomain, nil).Once()

		threadlike, err := threadlikeService.LikeController(context.Background(), threadlikeDomain, 1)
		assert.Nil(t, err)
		assert.Equal(t, threadlikeDomain, threadlike)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{User_id: 0}, nil).Once()
		threadlikeRepository.On("NewLike", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadlikeDomain, 1, errors.New("")).Once()
		upMockRepository.On("AddReputationPoint", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(userpointDomain, nil).Once()

		_, err := threadlikeService.LikeController(context.Background(), threadlikeDomain, threadlikeDomain.User_id)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Valid", func(t *testing.T) {
		threadlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{User_id: 1, State: false}, nil).Once()
		threadlikeRepository.On("Like", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadlikeDomain, 1, nil).Once()
		upMockRepository.On("AddReputationPoint", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(userpointDomain, nil).Once()

		threadlike, err := threadlikeService.LikeController(context.Background(), threadlikeDomain, threadlikeDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, threadlike.User_id)
	})
	t.Run("Test Case 4 | Invalid", func(t *testing.T) {
		threadlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{User_id: 1, State: false}, nil).Once()
		threadlikeRepository.On("Like", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadlikeDomain, 1, errors.New("")).Once()
		upMockRepository.On("AddReputationPoint", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(userpointDomain, nil).Once()

		_, err := threadlikeService.LikeController(context.Background(), threadlikeDomain, threadlikeDomain.User_id)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 5 | Valid", func(t *testing.T) {
		threadlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{User_id: 1, State: true}, nil).Once()
		threadlikeRepository.On("Unlike", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadlikeDomain, 1, nil).Once()
		upMockRepository.On("AddReputationPoint", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(userpointDomain, nil).Once()

		threadlike, err := threadlikeService.LikeController(context.Background(), threadlikeDomain, threadlikeDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, threadlike.User_id)
	})
	t.Run("Test Case 6 | Invalid", func(t *testing.T) {
		threadlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{User_id: 1, State: true}, nil).Once()
		threadlikeRepository.On("Unlike", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadlikeDomain, 1, errors.New("")).Once()
		upMockRepository.On("AddReputationPoint", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(userpointDomain, nil).Once()

		_, err := threadlikeService.LikeController(context.Background(), threadlikeDomain, threadlikeDomain.User_id)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 7 | Invalid", func(t *testing.T) {
		threadlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadlikeDomain, errors.New("")).Once()

		_, err := threadlikeService.LikeController(context.Background(), threadlikeDomain, threadlikeDomain.User_id)
		assert.NotNil(t, err)
	})
}
