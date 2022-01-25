package threadsaves_test

import (
	"context"
	"errors"
	ts "fgd-alterra-29/business/thread_saves"
	_threadsaveMocks "fgd-alterra-29/business/thread_saves/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var threadsaveRepository _threadsaveMocks.Repository
var threadsaveService ts.UseCase
var threadsaveDomain ts.Domain

func setup() {
	threadsaveService = ts.NewThreadSaveUseCase(&threadsaveRepository, time.Hour*1)
	threadsaveDomain = ts.Domain{
		User_id:  1,
		Saved_at: time.Time{},
	}
}

func TestGetSaveState(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadsaveRepository.On("GetSaveState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(ts.Domain{User_id: 0}, nil).Once()
		threadsaveRepository.On("NewSave", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadsaveDomain, nil).Once()

		threadsave, err := threadsaveService.SaveThreadController(context.Background(), threadsaveDomain, threadsaveDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, threadsave.User_id)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadsaveRepository.On("GetSaveState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(ts.Domain{User_id: 0}, nil).Once()
		threadsaveRepository.On("NewSave", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadsaveDomain, errors.New("")).Once()

		_, err := threadsaveService.SaveThreadController(context.Background(), threadsaveDomain, threadsaveDomain.User_id)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 3 | Valid", func(t *testing.T) {
		threadsaveRepository.On("GetSaveState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(ts.Domain{User_id: 1, State: false}, nil).Once()
		threadsaveRepository.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadsaveDomain, nil).Once()

		threadsave, err := threadsaveService.SaveThreadController(context.Background(), threadsaveDomain, threadsaveDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, threadsave.User_id)
	})
	t.Run("Test Case 4 | Invalid", func(t *testing.T) {
		threadsaveRepository.On("GetSaveState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(ts.Domain{User_id: 1, State: false}, nil).Once()
		threadsaveRepository.On("Save", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadsaveDomain, errors.New("")).Once()

		_, err := threadsaveService.SaveThreadController(context.Background(), threadsaveDomain, threadsaveDomain.User_id)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 5 | Valid", func(t *testing.T) {
		threadsaveRepository.On("GetSaveState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(ts.Domain{User_id: 1, State: true}, nil).Once()
		threadsaveRepository.On("Unsave", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadsaveDomain, nil).Once()

		threadsave, err := threadsaveService.SaveThreadController(context.Background(), threadsaveDomain, threadsaveDomain.User_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, threadsave.User_id)
	})
	t.Run("Test Case 6 | Invalid", func(t *testing.T) {
		threadsaveRepository.On("GetSaveState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(ts.Domain{User_id: 1, State: true}, nil).Once()
		threadsaveRepository.On("Unsave", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadsaveDomain, errors.New("")).Once()

		_, err := threadsaveService.SaveThreadController(context.Background(), threadsaveDomain, threadsaveDomain.User_id)
		assert.NotNil(t, err)
	})
	t.Run("Test Case 7 | Invalid", func(t *testing.T) {
		threadsaveRepository.On("GetSaveState", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadsaveDomain, errors.New("")).Once()

		_, err := threadsaveService.SaveThreadController(context.Background(), threadsaveDomain, threadsaveDomain.User_id)
		assert.NotNil(t, err)
	})
}
