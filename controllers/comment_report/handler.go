package commentreport

import (
	commentreport "fgd-alterra-29/business/comment_report"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/comment_report/request"
	"fgd-alterra-29/controllers/comment_report/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentReportController struct {
	CommentReportUseCase commentreport.UseCase
}

func NewCommentReportController(commentUseCase commentreport.UseCase) *CommentReportController {
	return &CommentReportController{
		CommentReportUseCase: commentUseCase,
	}
}

func (handler CommentReportController) GetReportsByCategoryController(c echo.Context) error {
	category := c.Param("category")
	ctx := c.Request().Context()

	commentreport, err := handler.CommentReportUseCase.SearchReportsByCategoryController(ctx, category)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, commentreport)
}

func (handler CommentReportController) GetCommentReportStat(c echo.Context) error {
	ctx := c.Request().Context()

	commentreport, err := handler.CommentReportUseCase.GetCommentReportStat(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, commentreport)
}

func (handler CommentReportController) CreateReportComment(c echo.Context) error {
	NewReport := request.CreateReport{}
	c.Bind(&NewReport)

	domain := NewReport.ToDomain()

	ctx := c.Request().Context()

	commentreport, err := handler.CommentReportUseCase.CreateReportComment(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, commentreport)
}

func (handler CommentReportController) AdminGetReports(c echo.Context) error {
	ctx := c.Request().Context()

	commentreport, err := handler.CommentReportUseCase.AdminGetReports(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListReports(commentreport))
}

func (handler CommentReportController) DeleteCommentReport(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	_, err := handler.CommentReportUseCase.DeleteCommentReport(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, "Report deleted success")
}
