package commentreport_test

import (
	"context"
	"errors"
	tr "fgd-alterra-29/business/comment_report"
	_commentreportMocks "fgd-alterra-29/business/comment_report/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var commentreportRepository _commentreportMocks.Repository
var commentreportService tr.UseCase
var commentreportDomain tr.Domain

func setup() {
	commentreportService = tr.NewCommentReportUseCase(&commentreportRepository, time.Hour*1)
	commentreportDomain = tr.Domain{
		Reporter_id: 1,
	}
}

func TestSearchReportsByCategory(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentreportRepository.On("SearchReportsByCategory", mock.Anything, mock.AnythingOfType("string")).Return([]tr.Domain{commentreportDomain}, nil).Once()

		commentreport, err := commentreportService.SearchReportsByCategoryController(context.Background(), "Spam")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(commentreport))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentreportRepository.On("SearchReportsByCategory", mock.Anything, mock.AnythingOfType("string")).Return([]tr.Domain{commentreportDomain}, errors.New("")).Once()

		_, err := commentreportService.SearchReportsByCategoryController(context.Background(), "Spam")
		assert.NotNil(t, err)
	})
}
func TestGetCommentReportStat(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentreportRepository.On("GetCommentReportStat", mock.Anything).Return([]tr.Domain{commentreportDomain}, nil).Once()

		commentreport, err := commentreportService.GetCommentReportStat(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(commentreport))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentreportRepository.On("GetCommentReportStat", mock.Anything).Return([]tr.Domain{commentreportDomain}, errors.New("")).Once()

		_, err := commentreportService.GetCommentReportStat(context.Background())
		assert.NotNil(t, err)
	})
}
func TestCreateReportComment(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentreportRepository.On("CreateReportComment", mock.Anything, mock.Anything).Return(commentreportDomain, nil).Once()

		commentreport, err := commentreportService.CreateReportComment(context.Background(), commentreportDomain)
		assert.Nil(t, err)
		assert.Equal(t, 1, commentreport.Reporter_id)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentreportRepository.On("CreateReportComment", mock.Anything, mock.Anything).Return(commentreportDomain, errors.New("")).Once()

		_, err := commentreportService.CreateReportComment(context.Background(), commentreportDomain)
		assert.NotNil(t, err)
	})
}
func TestAdminGetReports(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentreportRepository.On("AdminGetReports", mock.Anything).Return([]tr.Domain{commentreportDomain}, nil).Once()

		commentreport, err := commentreportService.AdminGetReports(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(commentreport))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentreportRepository.On("AdminGetReports", mock.Anything).Return([]tr.Domain{commentreportDomain}, errors.New("")).Once()

		_, err := commentreportService.AdminGetReports(context.Background())
		assert.NotNil(t, err)
	})
}
func TestDeleteCommentReport(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		commentreportRepository.On("DeleteCommentReport", mock.Anything, mock.AnythingOfType("int")).Return(commentreportDomain, nil).Once()

		commentreport, err := commentreportService.DeleteCommentReport(context.Background(), commentreportDomain.Reporter_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, commentreport.Reporter_id)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		commentreportRepository.On("DeleteCommentReport", mock.Anything, mock.AnythingOfType("int")).Return(commentreportDomain, errors.New("")).Once()

		_, err := commentreportService.DeleteCommentReport(context.Background(), commentreportDomain.Reporter_id)
		assert.NotNil(t, err)
	})
}
