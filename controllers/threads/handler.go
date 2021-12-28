package threads

import (
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/threads/request"
	"fgd-alterra-29/controllers/threads/responses"
	"net/http"
	"strconv"

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
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetProfileThreads(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListProfileThread(thread))
}

func (handler ThreadController) CreateThread(c echo.Context) error {
	NewThread := request.CreateThread{}
	c.Bind(&NewThread)

	domain := NewThread.ToDomain()

	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.CreateThread(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, thread)
}
