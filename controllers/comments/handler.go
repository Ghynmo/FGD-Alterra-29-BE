package comments

import (
	"fgd-alterra-29/business/comments"
	"fgd-alterra-29/controllers"
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

func (handler CommentController) GetPostsByCommentController(c echo.Context) error {
	comment := c.Param("comment")
	ctx := c.Request().Context()

	comments, err := handler.CommentUseCase.GetPostsByCommentController(ctx, comment)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListPostProfile(comments))
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

func (handler CommentController) GetPostsController(c echo.Context) error {
	ctx := c.Request().Context()

	posts, err := handler.CommentUseCase.GetPosts(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListPosts(posts))
}

func (handler CommentController) DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	_, err := handler.CommentUseCase.DeletePost(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Post deleted success")
}
