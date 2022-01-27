package follows_test

import (
	"context"
	"errors"
	"fgd-alterra-29/business/follows"
	_followMocks "fgd-alterra-29/business/follows/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var followRepository _followMocks.Repository
var followService follows.UseCase
var followDomain follows.Domain

func setup() {
	followService = follows.NewFollowUseCase(&followRepository, time.Hour*1)
	followDomain = follows.Domain{
		User_id:     1,
		Follower_id: 2,
	}
}

func TestGetFollowers(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		followRepository.On("GetFollowers", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]follows.Domain{followDomain}, nil).Once()

		follow, err := followService.GetFollowers(context.Background(), followDomain.User_id, followDomain.Follower_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(follow))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		followRepository.On("GetFollowers", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]follows.Domain{followDomain}, errors.New("")).Once()

		_, err := followService.GetFollowers(context.Background(), followDomain.User_id, followDomain.Follower_id)
		assert.NotNil(t, err)
	})
}
func TestGetFollowing(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		followRepository.On("GetFollowing", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]follows.Domain{followDomain}, nil).Once()

		follow, err := followService.GetFollowing(context.Background(), followDomain.User_id, followDomain.Follower_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(follow))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		followRepository.On("GetFollowing", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]follows.Domain{followDomain}, errors.New("")).Once()

		_, err := followService.GetFollowing(context.Background(), followDomain.User_id, followDomain.Follower_id)
		assert.NotNil(t, err)
	})
}
func TestGetFollowState(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		followRepository.On("GetFollowState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(follows.Domain{User_id: 0}, nil).Once()
		followRepository.On("NewFollow", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(followDomain, nil).Once()

		follow, err := followService.FollowsController(context.Background(), followDomain, followDomain.Follower_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, follow.User_id)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		followRepository.On("GetFollowState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(follows.Domain{User_id: 0}, nil).Once()
		followRepository.On("NewFollow", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(followDomain, errors.New("")).Once()

		_, err := followService.FollowsController(context.Background(), followDomain, followDomain.Follower_id)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Valid", func(t *testing.T) {
		followRepository.On("GetFollowState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(follows.Domain{User_id: 1, State: false}, nil).Once()
		followRepository.On("Follows", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(followDomain, nil).Once()

		follow, err := followService.FollowsController(context.Background(), followDomain, followDomain.Follower_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, follow.User_id)
	})
	t.Run("Test Case 4 | Invalid", func(t *testing.T) {
		followRepository.On("GetFollowState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(follows.Domain{User_id: 1, State: false}, nil).Once()
		followRepository.On("Follows", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(followDomain, errors.New("")).Once()

		_, err := followService.FollowsController(context.Background(), followDomain, followDomain.Follower_id)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 5 | Valid", func(t *testing.T) {
		followRepository.On("GetFollowState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(follows.Domain{User_id: 1, State: true}, nil).Once()
		followRepository.On("Unfollow", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(followDomain, nil).Once()

		follow, err := followService.FollowsController(context.Background(), followDomain, followDomain.Follower_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, follow.User_id)
	})
	t.Run("Test Case 6 | Invalid", func(t *testing.T) {
		followRepository.On("GetFollowState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(follows.Domain{User_id: 1, State: true}, nil).Once()
		followRepository.On("Unfollow", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(followDomain, errors.New("")).Once()

		_, err := followService.FollowsController(context.Background(), followDomain, followDomain.Follower_id)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 7 | Invalid", func(t *testing.T) {
		followRepository.On("GetFollowState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(followDomain, errors.New("CANNOT FOLLOWS YOURSELF")).Once()

		_, err := followService.FollowsController(context.Background(), followDomain, followDomain.Follower_id)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 8 | Invalid", func(t *testing.T) {
		followDomain.User_id = 1
		followDomain.Follower_id = 1
		_, err := followService.FollowsController(context.Background(), followDomain, followDomain.Follower_id)
		assert.NotNil(t, err)
	})
}
