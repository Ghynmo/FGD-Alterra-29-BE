package threadlikes

import (
	threadlikes "fgd-alterra-29/business/thread_likes"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/thread_likes/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ThreadLikeController struct {
	ThreadLikeUseCase threadlikes.UseCase
}

func NewThreadLikeController(threadUseCase threadlikes.UseCase) *ThreadLikeController {
	return &ThreadLikeController{
		ThreadLikeUseCase: threadUseCase,
	}
}

func (handler ThreadLikeController) CreateLikes(c echo.Context) error {
	NewLike := request.Like{}
	c.Bind(&NewLike)

	domain := NewLike.ToDomain()

	ctx := c.Request().Context()

	threadlikes, err := handler.ThreadLikeUseCase.LikeController(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, threadlikes)
}

func (handler ThreadLikeController) DeleteLikes(c echo.Context) error {
	NewUnlike := request.Like{}
	c.Bind(&NewUnlike)

	domain := NewUnlike.ToDomain()

	ctx := c.Request().Context()

	threadlikes, err := handler.ThreadLikeUseCase.UnlikeController(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, threadlikes)
}
