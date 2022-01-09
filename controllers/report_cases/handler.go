package reportcases

import (
	reportcases "fgd-alterra-29/business/report_cases"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/report_cases/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReportCaseController struct {
	ReportCaseUseCase reportcases.UseCase
}

func NewReportCaseController(reportcasesUseCase reportcases.UseCase) *ReportCaseController {
	return &ReportCaseController{
		ReportCaseUseCase: reportcasesUseCase,
	}
}

func (handler ReportCaseController) GetReportForm(c echo.Context) error {
	ctx := c.Request().Context()

	reportcases, err := handler.ReportCaseUseCase.GetReportForm(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListReportForm(reportcases))
}
