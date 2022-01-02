package threads

import (
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/controllers"
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

func (handler ThreadController) GetHomepageThreads(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetHomepageThreads(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListRecommendationThreads(thread))
}

func (handler ThreadController) GetRecommendationThreads(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetRecommendationThreads(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListRecommendationThreads(thread))
}

func (handler ThreadController) GetHotThreads(c echo.Context) error {
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetHotThreads(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListRecommendationThreads(thread))
}
