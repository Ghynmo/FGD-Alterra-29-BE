package threads

import (
	"fgd-alterra-29/business/comments"
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/threads/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ThreadController struct {
	ThreadUseCase  threads.UseCase
	CommentUseCase comments.UseCase
}

func NewThreadController(threadUseCase threads.UseCase, commentUseCase comments.UseCase) *ThreadController {
	return &ThreadController{
		ThreadUseCase:  threadUseCase,
		CommentUseCase: commentUseCase,
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

func (handler ThreadController) GetRecommendationThreads(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetRecommendationThreads(ctx, id)
	_, err1 := handler.CommentUseCase.GetCommentByThread(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err1 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	fmt.Println("Thread in Ctrl", thread)
	return controllers.NewSuccessResponse(c, responses.ToListRecommendationThreads(thread))
}

func (handler ThreadController) GetHotThreads(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	thread, err := handler.ThreadUseCase.GetHotThreads(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToListRecommendationThreads(thread))
}
