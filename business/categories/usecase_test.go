package categories_test

import (
	"context"
	"errors"
	"fgd-alterra-29/business/categories"
	_categoryMocks "fgd-alterra-29/business/categories/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository _categoryMocks.Repository
var categoryService categories.UseCase
var categoryDomain categories.Domain

func setup() {
	categoryService = categories.NewCategoryUseCase(&categoryRepository, time.Hour*1)
	categoryDomain = categories.Domain{
		ID:       1,
		Category: "Social",
	}
}

func TestGetCategories(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		categoryRepository.On("GetCategories", mock.Anything).Return([]categories.Domain{categoryDomain}, nil).Once()

		category, err := categoryService.GetCategoriesController(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(category))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		categoryRepository.On("GetCategories", mock.Anything).Return([]categories.Domain{categoryDomain}, errors.New("")).Once()

		_, err := categoryService.GetCategoriesController(context.Background())
		assert.NotNil(t, err)
	})
}
func TestGetUserActiveInCategory(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		categoryRepository.On("GetUserActiveInCategory", mock.Anything, mock.AnythingOfType("int")).Return([]categories.Domain{categoryDomain}, nil).Once()

		category, err := categoryService.GetUserActiveInCategory(context.Background(), categoryDomain.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(category))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		categoryRepository.On("GetUserActiveInCategory", mock.Anything, mock.AnythingOfType("int")).Return([]categories.Domain{categoryDomain}, errors.New("")).Once()

		_, err := categoryService.GetUserActiveInCategory(context.Background(), categoryDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestCreateCategories(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		categoryRepository.On("CreateCategories", mock.Anything, mock.Anything).Return(categoryDomain, nil).Once()

		category, err := categoryService.CreateCategoriesController(context.Background(), categoryDomain)
		assert.Nil(t, err)
		assert.Equal(t, 1, category.ID)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		categoryRepository.On("CreateCategories", mock.Anything, mock.Anything).Return(categoryDomain, errors.New("")).Once()

		_, err := categoryService.CreateCategoriesController(context.Background(), categoryDomain)
		assert.NotNil(t, err)
	})
}
