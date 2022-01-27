package threadreport_test

import (
	"context"
	"errors"
	tr "fgd-alterra-29/business/thread_report"
	_threadreportMocks "fgd-alterra-29/business/thread_report/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var threadreportRepository _threadreportMocks.Repository
var threadreportService tr.UseCase
var threadreportDomain tr.Domain

func setup() {
	threadreportService = tr.NewThreadReportUseCase(&threadreportRepository, time.Hour*1)
	threadreportDomain = tr.Domain{
		Reporter_id: 1,
	}
}

func TestSearchReportsByCategory(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadreportRepository.On("SearchReportsByCategory", mock.Anything, mock.AnythingOfType("string")).Return([]tr.Domain{threadreportDomain}, nil).Once()

		threadreport, err := threadreportService.SearchReportsByCategoryController(context.Background(), "Spam")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(threadreport))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadreportRepository.On("SearchReportsByCategory", mock.Anything, mock.AnythingOfType("string")).Return([]tr.Domain{threadreportDomain}, errors.New("")).Once()

		_, err := threadreportService.SearchReportsByCategoryController(context.Background(), "Spam")
		assert.NotNil(t, err)
	})
}
func TestGetThreadReportStat(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadreportRepository.On("GetThreadReportStat", mock.Anything).Return([]tr.Domain{threadreportDomain}, nil).Once()

		threadreport, err := threadreportService.GetThreadReportStat(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(threadreport))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadreportRepository.On("GetThreadReportStat", mock.Anything).Return([]tr.Domain{threadreportDomain}, errors.New("")).Once()

		_, err := threadreportService.GetThreadReportStat(context.Background())
		assert.NotNil(t, err)
	})
}
func TestCreateReportThread(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadreportRepository.On("CreateReportThread", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadreportDomain, nil).Once()

		threadreport, err := threadreportService.CreateReportThread(context.Background(), threadreportDomain, threadreportDomain.Reporter_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, threadreport.Reporter_id)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadreportRepository.On("CreateReportThread", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(threadreportDomain, errors.New("")).Once()

		_, err := threadreportService.CreateReportThread(context.Background(), threadreportDomain, threadreportDomain.Reporter_id)
		assert.NotNil(t, err)
	})
}
func TestAdminGetReports(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadreportRepository.On("AdminGetReports", mock.Anything).Return([]tr.Domain{threadreportDomain}, nil).Once()

		threadreport, err := threadreportService.AdminGetReports(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(threadreport))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadreportRepository.On("AdminGetReports", mock.Anything).Return([]tr.Domain{threadreportDomain}, errors.New("")).Once()

		_, err := threadreportService.AdminGetReports(context.Background())
		assert.NotNil(t, err)
	})
}
func TestSolvedThreadReport(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		threadreportRepository.On("SolvedThreadReport", mock.Anything, mock.AnythingOfType("int")).Return(threadreportDomain, nil).Once()

		threadreport, err := threadreportService.SolvedThreadReport(context.Background(), threadreportDomain.Reporter_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, threadreport.Reporter_id)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		threadreportRepository.On("SolvedThreadReport", mock.Anything, mock.AnythingOfType("int")).Return(threadreportDomain, errors.New("")).Once()

		_, err := threadreportService.SolvedThreadReport(context.Background(), threadreportDomain.Reporter_id)
		assert.NotNil(t, err)
	})
}
