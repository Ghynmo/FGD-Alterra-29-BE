package comments

import (
	"fgd-alterra-29/business/comments"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/comments/request"
	"fgd-alterra-29/controllers/comments/responses"
	"net/http"
	"strconv"

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
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	comments, err := handler.CommentUseCase.GetCommentProfile(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListPostProfile(comments))
}

func (handler CommentController) CreateCommentController(c echo.Context) error {
	NewComment := request.CreateComment{}
	c.Bind(&NewComment)

	domain := NewComment.ToDomain()

	ctx := c.Request().Context()

	comments, err := handler.CommentUseCase.CreateCommentController(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, comments)
}
