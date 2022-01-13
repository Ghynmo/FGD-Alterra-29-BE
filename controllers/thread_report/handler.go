package threadreport

import (
	threadreport "fgd-alterra-29/business/thread_report"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/thread_report/request"
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

func (handler ThreadReportController) SearchReportsByCategoryController(c echo.Context) error {
	category := c.Param("category")
	ctx := c.Request().Context()

	threadreport, err := handler.ThreadReportUseCase.SearchReportsByCategoryController(ctx, category)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, threadreport)
}

func (handler ThreadReportController) GetThreadReportStat(c echo.Context) error {
	ctx := c.Request().Context()

	threadreport, err := handler.ThreadReportUseCase.GetThreadReportStat(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, threadreport)
}

func (handler ThreadReportController) CreateReportThread(c echo.Context) error {
	NewReport := request.CreateReport{}
	c.Bind(&NewReport)

	domain := NewReport.ToDomain()

	ctx := c.Request().Context()

	threadreport, err := handler.ThreadReportUseCase.CreateReportThread(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, threadreport)
}

func (handler ThreadReportController) AdminGetReports(c echo.Context) error {
	ctx := c.Request().Context()

	threadreport, err := handler.ThreadReportUseCase.AdminGetReports(ctx)
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
