package commentlikes_test

// import (
// 	"context"
// 	"errors"
// 	tl "fgd-alterra-29/business/comment_likes"
// 	_commentlikeMocks "fgd-alterra-29/business/comment_likes/mocks"
// 	userpoint "fgd-alterra-29/business/user_points"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// var commentlikeRepository _commentlikeMocks.Repository
// var userpointRepository userpoint.Repository
// var commentlikeService tl.UseCase
// var commentlikeDomain tl.Domain

// func setup() {
// 	commentlikeService = tl.NewCommentLikeUseCase(&commentlikeRepository, time.Hour*1, userpointRepository)
// 	commentlikeDomain = tl.Domain{
// 		Liker_id: 1,
// 		Liked_at: time.Time{},
// 	}
// }

// func TestGetLikeState(t *testing.T) {
// 	setup()
// 	t.Run("Test Case 1 | Valid", func(t *testing.T) {
// 		commentlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{Liker_id: 0}, nil).Once()
// 		commentlikeRepository.On("NewLike", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(commentlikeDomain, nil).Once()

// 		commentlike, err := commentlikeService.LikeController(context.Background(), commentlikeDomain, commentlikeDomain.Liker_id)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, commentlike.Liker_id)
// 	})
// 	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
// 		commentlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{Liker_id: 0}, nil).Once()
// 		commentlikeRepository.On("NewLike", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(commentlikeDomain, errors.New("")).Once()

// 		_, err := commentlikeService.LikeController(context.Background(), commentlikeDomain, commentlikeDomain.Liker_id)
// 		assert.NotNil(t, err)
// 	})
// 	t.Run("Test Case 3 | Valid", func(t *testing.T) {
// 		commentlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{Liker_id: 1, State: false}, nil).Once()
// 		commentlikeRepository.On("Like", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(commentlikeDomain, nil).Once()

// 		commentlike, err := commentlikeService.LikeController(context.Background(), commentlikeDomain, commentlikeDomain.Liker_id)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, commentlike.Liker_id)
// 	})
// 	t.Run("Test Case 4 | Invalid", func(t *testing.T) {
// 		commentlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{Liker_id: 1, State: false}, nil).Once()
// 		commentlikeRepository.On("Like", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(commentlikeDomain, errors.New("")).Once()

// 		_, err := commentlikeService.LikeController(context.Background(), commentlikeDomain, commentlikeDomain.Liker_id)
// 		assert.NotNil(t, err)
// 	})
// 	t.Run("Test Case 5 | Valid", func(t *testing.T) {
// 		commentlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{Liker_id: 1, State: true}, nil).Once()
// 		commentlikeRepository.On("Unlike", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(commentlikeDomain, nil).Once()

// 		commentlike, err := commentlikeService.LikeController(context.Background(), commentlikeDomain, commentlikeDomain.Liker_id)
// 		assert.Nil(t, err)
// 		assert.Equal(t, 1, commentlike.Liker_id)
// 	})
// 	t.Run("Test Case 6 | Invalid", func(t *testing.T) {
// 		commentlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(tl.Domain{Liker_id: 1, State: true}, nil).Once()
// 		commentlikeRepository.On("Unlike", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(commentlikeDomain, errors.New("")).Once()

// 		_, err := commentlikeService.LikeController(context.Background(), commentlikeDomain, commentlikeDomain.Liker_id)
// 		assert.NotNil(t, err)
// 	})
// 	t.Run("Test Case 7 | Invalid", func(t *testing.T) {
// 		commentlikeRepository.On("GetLikeState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(commentlikeDomain, errors.New("")).Once()

// 		_, err := commentlikeService.LikeController(context.Background(), commentlikeDomain, commentlikeDomain.Liker_id)
// 		assert.NotNil(t, err)
// 	})
// }
