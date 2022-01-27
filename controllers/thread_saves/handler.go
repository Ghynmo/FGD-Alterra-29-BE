package threadsaves

import (
	"fgd-alterra-29/app/middlewares"
	threadsaves "fgd-alterra-29/business/thread_saves"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/thread_saves/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ThreadSaveController struct {
	ThreadSaveUseCase threadsaves.UseCase
}

func NewThreadSaveController(threadUseCase threadsaves.UseCase) *ThreadSaveController {
	return &ThreadSaveController{
		ThreadSaveUseCase: threadUseCase,
	}
}

func (handler ThreadSaveController) SaveThread(c echo.Context) error {
	id := middlewares.ExtractID(c)

	NewSave := request.Save{}
	c.Bind(&NewSave)

	domain := NewSave.ToDomain()

	ctx := c.Request().Context()

	threadsaves, err := handler.ThreadSaveUseCase.SaveThreadController(ctx, domain, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, threadsaves)
}
