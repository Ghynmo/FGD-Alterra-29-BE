package threadlikes

import (
	"fgd-alterra-29/app/middlewares"
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

func (handler ThreadLikeController) Likes(c echo.Context) error {
	NewLike := request.Like{}
	c.Bind(&NewLike)

	domain := NewLike.ToDomain()
	id := middlewares.ExtractID(c)

	ctx := c.Request().Context()

	threadlikes, err := handler.ThreadLikeUseCase.LikeController(ctx, domain, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, threadlikes)
}
