package comments

import (
	"fgd-alterra-29/business/comments"
	"fgd-alterra-29/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	CommentUseCase comments.UseCase
}

func NewCommentController(threadUseCase comments.UseCase) *CommentController {
	return &CommentController{
		CommentUseCase: threadUseCase,
	}
}

func (handler CommentController) GetProfileComments(c echo.Context) error {
	ctx := c.Request().Context()
	id := 1

	comments, err := handler.CommentUseCase.GetCommentProfile(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, comments)
}
