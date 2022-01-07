package threadreport

import (
	threadreport "fgd-alterra-29/business/thread_report"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/thread_report/responses"
	"net/http"
	"strconv"

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

func (handler ThreadReportController) GetReportsByCategoryController(c echo.Context) error {
	category := c.Param("category")
	ctx := c.Request().Context()

	threadreport, err := handler.ThreadReportUseCase.GetReportsByCategoryController(ctx, category)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, threadreport)
}

func (handler ThreadReportController) GetThreadReports(c echo.Context) error {
	ctx := c.Request().Context()

	threadreport, err := handler.ThreadReportUseCase.GetThreadReports(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, threadreport)
}

func (handler ThreadReportController) GetReports(c echo.Context) error {
	ctx := c.Request().Context()

	threadreport, err := handler.ThreadReportUseCase.GetReports(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListReports(threadreport))
}

func (handler ThreadReportController) DeleteThreadReport(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	_, err := handler.ThreadReportUseCase.DeleteThreadReport(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, "Report deleted success")
}
