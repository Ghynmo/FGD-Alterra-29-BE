package threadreport

import (
	threadreport "fgd-alterra-29/business/thread_report"
	"fgd-alterra-29/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ThreadReportController struct {
	ThreadReportUseCase threadreport.UseCase
}

func NewThreadReportController(threadUseCase threadreport.UseCase) *ThreadReportController {
	return &ThreadReportController{
		ThreadReportUseCase: threadUseCase,
	}
}

func (handler ThreadReportController) GetThreadReports(c echo.Context) error {
	ctx := c.Request().Context()

	threadreport, err := handler.ThreadReportUseCase.GetThreadReports(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, threadreport)
}
