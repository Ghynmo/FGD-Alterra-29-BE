package reportcases_test

import (
	"context"
	"errors"
	rc "fgd-alterra-29/business/report_cases"
	_reportcaseMocks "fgd-alterra-29/business/report_cases/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var reportcaseRepository _reportcaseMocks.Repository
var reportcaseService rc.UseCase
var reportcaseDomain rc.Domain

func setup() {
	reportcaseService = rc.NewReportCaseUseCase(&reportcaseRepository, time.Hour*1)
	reportcaseDomain = rc.Domain{
		ID: 1,
	}
}

func TestGetReportForm(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		reportcaseRepository.On("GetReportForm", mock.Anything).Return([]rc.Domain{reportcaseDomain}, nil).Once()

		reportcase, err := reportcaseService.GetReportForm(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, 1, len(reportcase))
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		reportcaseRepository.On("GetReportForm", mock.Anything).Return([]rc.Domain{reportcaseDomain}, errors.New("")).Once()

		_, err := reportcaseService.GetReportForm(context.Background())
		assert.NotNil(t, err)
	})
}
func TestCreateCaseController(t *testing.T) {
	setup()
	t.Run("Test Case 1 | Valid", func(t *testing.T) {
		reportcaseRepository.On("CreateCase", mock.Anything, mock.Anything).Return(reportcaseDomain, nil).Once()

		reportcase, err := reportcaseService.CreateCaseController(context.Background(), reportcaseDomain)
		assert.Nil(t, err)
		assert.Equal(t, 1, reportcase.ID)
	})
	t.Run("Test Case 2 | Invalid", func(t *testing.T) {
		reportcaseRepository.On("CreateCase", mock.Anything, mock.Anything).Return(reportcaseDomain, errors.New("")).Once()

		_, err := reportcaseService.CreateCaseController(context.Background(), reportcaseDomain)
		assert.NotNil(t, err)
	})
}
