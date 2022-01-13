package commentlikes

import (
	commentlikes "fgd-alterra-29/business/comment_likes"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/comment_likes/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentLikeController struct {
	CommentLikeUseCase commentlikes.UseCase
}

func NewCommentLikeController(threadUseCase commentlikes.UseCase) *CommentLikeController {
	return &CommentLikeController{
		CommentLikeUseCase: threadUseCase,
	}
}

func (handler CommentLikeController) Likes(c echo.Context) error {
	NewLike := request.Like{}
	c.Bind(&NewLike)

	domain := NewLike.ToDomain()

	ctx := c.Request().Context()

	commentlikes, err := handler.CommentLikeUseCase.LikeController(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NoDataSuccessResponse(c, commentlikes)
}
