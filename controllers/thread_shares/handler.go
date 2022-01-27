package threadshares

import (
	"fgd-alterra-29/app/middlewares"
	threadshares "fgd-alterra-29/business/thread_shares"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/thread_shares/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ThreadShareController struct {
	ThreadShareUseCase threadshares.UseCase
}

func NewThreadShareController(threadUseCase threadshares.UseCase) *ThreadShareController {
	return &ThreadShareController{
		ThreadShareUseCase: threadUseCase,
	}
}

func (handler ThreadShareController) ShareThread(c echo.Context) error {
	id := middlewares.ExtractID(c)
	NewShare := request.Share{}
	c.Bind(&NewShare)

	domain := NewShare.ToDomain()

	ctx := c.Request().Context()

	threadshares, err := handler.ThreadShareUseCase.ThreadShareController(ctx, domain, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, threadshares)
}
