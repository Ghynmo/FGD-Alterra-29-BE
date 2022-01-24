package comments

import (
	"fgd-alterra-29/app/middlewares"
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

func (handler CommentController) GetPostsByCommentController(c echo.Context) error {
	comment := c.Param("comment")
	ctx := c.Request().Context()

	comments, err := handler.CommentUseCase.GetPostsByCommentController(ctx, comment)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListPosts(comments))
}

func (handler CommentController) GetCommentByThreadController(c echo.Context) error {
	id := middlewares.ExtractID(c)

	GetComment := request.GetByThread{}
	c.Bind(&GetComment)

	thread_id := GetComment.Thread_id

	ctx := c.Request().Context()

	comments, err := handler.CommentUseCase.GetCommentByThreadController(ctx, thread_id, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListCommentRecommendation(comments))
}

func (handler CommentController) GetReplyComments(c echo.Context) error {
	id := middlewares.ExtractID(c)

	reply_of, _ := strconv.Atoi(c.Param("reply_of"))
	ctx := c.Request().Context()

	comments, err := handler.CommentUseCase.GetCommentReply(ctx, id, reply_of)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListCommentRecommendation(comments))
}

func (handler CommentController) GetProfileCommentsController(c echo.Context) error {
	id := middlewares.ExtractID(c)
	ctx := c.Request().Context()

	comments, err := handler.CommentUseCase.GetCommentProfileController(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListPostProfile(comments))
}

func (handler CommentController) GetProfileComments(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	comments, err := handler.CommentUseCase.GetCommentProfileController(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListPostProfile(comments))
}

func (handler CommentController) GetPostsController(c echo.Context) error {
	ctx := c.Request().Context()

	posts, err := handler.CommentUseCase.GetPostsController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ToListPosts(posts))
}

func (handler CommentController) UnactivatingPostController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	_, err := handler.CommentUseCase.UnactivatingPostController(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Post Unactivated")
}

func (handler CommentController) ActivatingPostController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	_, err := handler.CommentUseCase.ActivatingPostController(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, "Post Activated")
}

func (handler CommentController) CreateCommentController(c echo.Context) error {
	id := middlewares.ExtractID(c)
	NewComment := request.CreateComment{}
	c.Bind(&NewComment)

	domain := NewComment.ToDomain()

	ctx := c.Request().Context()

	comments, err := handler.CommentUseCase.CreateCommentController(ctx, domain, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NoDataSuccessResponse(c, comments)
}
