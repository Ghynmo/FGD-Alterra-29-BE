package threads

import (
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/controllers"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ThreadController struct {
	ThreadUseCase threads.UseCase
}

func NewThreadController(threadUseCase threads.UseCase) *ThreadController {
	return &ThreadController{
		ThreadUseCase: threadUseCase,
	}
}

func (handler ThreadController) GetProfileThreads(c echo.Context) error {
	ctx := c.Request().Context()
	id := 2

	thread, err := handler.ThreadUseCase.GetProfileThreads(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	fmt.Println(thread)
	return controllers.NewSuccessResponse(c, nil)
}
