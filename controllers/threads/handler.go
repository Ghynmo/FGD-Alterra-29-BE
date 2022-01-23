package threads

import (
	"fgd-alterra-29/app/middlewares"
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

func (handler ThreadController) GetThreadsByTitleController(c echo.Context) error {
	title := c.Param("title")
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetThreadsByTitleController(ctx, title)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListThread(thread))
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

func (handler ThreadController) GetThreadsController(c echo.Context) error {
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetThreads(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListThread(thread))
}

func (handler ThreadController) DeleteThread(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.DeleteThread(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.DeleteSuccessResponse(c, thread)
}

func (handler ThreadController) ActivateThread(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.ActivateThread(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, thread)
}

func (handler ThreadController) GetHomepageThreads(c echo.Context) error {
	id := middlewares.ExtractID(c)
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetHomepageThreads(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListHomeThreads(thread))
}

func (handler ThreadController) GetRecommendationThreads(c echo.Context) error {
	id := middlewares.ExtractID(c)
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetRecommendationThreads(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListHomeThreads(thread))
}

func (handler ThreadController) GetHotThreads(c echo.Context) error {
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetHotThreads(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListHomeThreads(thread))
}

func (handler ThreadController) GetSearch(c echo.Context) error {
	threadname := c.QueryParam("threadname")
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetSearch(ctx, threadname)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListHomeThreads(thread))
}

func (handler ThreadController) CreateThread(c echo.Context) error {
	id := middlewares.ExtractID(c)
	NewThread := request.CreateThread{}
	c.Bind(&NewThread)

	domain := NewThread.ToDomain()

	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.CreateThread(ctx, domain, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, thread)
}
