package comments_test

import (
	"context"
	"errors"
	"fgd-alterra-29/business/comments"
	_commentMocks "fgd-alterra-29/business/comments/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var commentRepository _commentMocks.Repository
var commentService comments.UseCase
var commentDomain comments.Domain

func setup() {
	commentService = comments.NewCommentUseCase(&commentRepository, time.Hour*1)
	commentDomain = comments.Domain{
		Comment:   "test hai",
		Thread_id: 1,
		User_id:   1,
		Q_Post:    1,
	}
}

func TestGetPostsByComment(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentRepository.On("GetPostsByComment", mock.Anything, mock.AnythingOfType("string")).Return([]comments.Domain{commentDomain}, nil).Once()

		comment, err := commentService.GetPostsByCommentController(context.Background(), commentDomain.Comment)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(comment))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentRepository.On("GetPostsByComment", mock.Anything, mock.AnythingOfType("string")).Return([]comments.Domain{commentDomain}, errors.New("")).Once()

		_, err := commentService.GetPostsByCommentController(context.Background(), commentDomain.Comment)
		assert.NotNil(t, err)
	})
}
func TestGetCommentByThread(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentRepository.On("GetCommentByThread", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]comments.Domain{commentDomain}, nil).Once()

		comment, err := commentService.GetCommentByThreadController(context.Background(), commentDomain.Thread_id, commentDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(comment))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentRepository.On("GetCommentByThread", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]comments.Domain{commentDomain}, errors.New("")).Once()

		_, err := commentService.GetCommentByThreadController(context.Background(), commentDomain.Thread_id, commentDomain.User_id)
		assert.NotNil(t, err)
	})
}
func TestGetCommentReply(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentRepository.On("GetCommentReply", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]comments.Domain{commentDomain}, nil).Once()

		comment, err := commentService.GetCommentReply(context.Background(), commentDomain.User_id, commentDomain.ReplyOf)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(comment))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentRepository.On("GetCommentReply", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]comments.Domain{commentDomain}, errors.New("")).Once()

		_, err := commentService.GetCommentReply(context.Background(), commentDomain.User_id, commentDomain.ReplyOf)
		assert.NotNil(t, err)
	})
}
func TestGetCommentProfile(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentRepository.On("GetCommentProfile", mock.Anything, mock.AnythingOfType("int")).Return([]comments.Domain{commentDomain}, nil).Once()

		comment, err := commentService.GetCommentProfileController(context.Background(), commentDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(comment))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentRepository.On("GetCommentProfile", mock.Anything, mock.AnythingOfType("int")).Return([]comments.Domain{commentDomain}, errors.New("")).Once()

		_, err := commentService.GetCommentProfileController(context.Background(), commentDomain.User_id)
		assert.NotNil(t, err)
	})
}
func TestGetPostQuantity(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentRepository.On("GetPostQuantity", mock.Anything).Return(commentDomain, nil).Once()

		comment, err := commentService.GetPostQuantityController(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, comment.Q_Post)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentRepository.On("GetPostQuantity", mock.Anything).Return(commentDomain, errors.New("")).Once()

		_, err := commentService.GetPostQuantityController(context.Background())
		assert.NotNil(t, err)
	})
}
func TestCreateComment(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentRepository.On("CreateComment", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(commentDomain, nil).Once()

		comment, err := commentService.CreateCommentController(context.Background(), commentDomain, commentDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, comment.Q_Post)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentRepository.On("CreateComment", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(commentDomain, errors.New("")).Once()

		_, err := commentService.CreateCommentController(context.Background(), commentDomain, commentDomain.User_id)
		assert.NotNil(t, err)
	})
}
func TestGetPosts(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentRepository.On("GetPosts", mock.Anything).Return([]comments.Domain{commentDomain}, nil).Once()

		comment, err := commentService.GetPostsController(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(comment))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentRepository.On("GetPosts", mock.Anything).Return([]comments.Domain{commentDomain}, errors.New("")).Once()

		_, err := commentService.GetPostsController(context.Background())
		assert.NotNil(t, err)
	})
}
func TestUnactivatingPost(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentRepository.On("UnactivatingPost", mock.Anything, mock.AnythingOfType("int")).Return(commentDomain, nil).Once()

		comment, err := commentService.UnactivatingPostController(context.Background(), commentDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, comment.Q_Post)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentRepository.On("UnactivatingPost", mock.Anything, mock.AnythingOfType("int")).Return(commentDomain, errors.New("")).Once()

		_, err := commentService.UnactivatingPostController(context.Background(), commentDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestActivatingPost(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentRepository.On("ActivatingPost", mock.Anything, mock.AnythingOfType("int")).Return(commentDomain, nil).Once()

		comment, err := commentService.ActivatingPostController(context.Background(), commentDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, comment.Q_Post)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentRepository.On("ActivatingPost", mock.Anything, mock.AnythingOfType("int")).Return(commentDomain, errors.New("")).Once()

		_, err := commentService.ActivatingPostController(context.Background(), commentDomain.ID)
		assert.NotNil(t, err)
	})
}
