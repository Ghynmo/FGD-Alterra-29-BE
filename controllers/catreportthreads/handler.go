package catreportthreads

import (
	"fgd-alterra-29/business/catreportthreads"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/catreportthreads/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CatReportThreadController struct {
	CatReportThreadUseCase catreportthreads.UseCase
}

func NewCatReportThreadController(catreportthreadUseCase catreportthreads.UseCase) *CatReportThreadController {
	return &CatReportThreadController{
		CatReportThreadUseCase: catreportthreadUseCase,
	}
}

func (handler CatReportThreadController) GetReportForm(c echo.Context) error {
	ctx := c.Request().Context()

	catreportthread, err := handler.CatReportThreadUseCase.GetReportForm(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListReportForm(catreportthread))
}
