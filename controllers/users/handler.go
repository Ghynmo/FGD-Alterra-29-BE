package users

import (
	"fgd-alterra-29/business/categories"
	"fgd-alterra-29/business/comments"
	threadreport "fgd-alterra-29/business/thread_report"
	"fgd-alterra-29/business/threads"
	userbadges "fgd-alterra-29/business/user_badges"
	"fgd-alterra-29/business/users"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/users/responses"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase         users.UseCase
	ThreadUseCase       threads.UseCase
	UserBadgeUseCase    userbadges.UseCase
	CategoryUseCase     categories.UseCase
	ThreadReportUseCase threadreport.UseCase
	CommentUseCase      comments.UseCase
}

func NewUserController(userUC users.UseCase, threadUC threads.UseCase, userbadgeUC userbadges.UseCase, categoryUC categories.UseCase, threadreportUC threadreport.UseCase, commentUC comments.UseCase) *UserController {
	return &UserController{
		UserUseCase:         userUC,
		ThreadUseCase:       threadUC,
		UserBadgeUseCase:    userbadgeUC,
		CategoryUseCase:     categoryUC,
		ThreadReportUseCase: threadreportUC,
		CommentUseCase:      commentUC,
	}
}

func (handler UserController) GetUsersController(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetUsersController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromListDomain(user))
}

func (handler UserController) GetProfileController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetProfileController(ctx, id)
	thread, err1 := handler.ThreadUseCase.GetProfileThreads(ctx, id)
	userbadges, err2 := handler.UserBadgeUseCase.GetUserBadge(ctx, id)
	catthreads, err3 := handler.CategoryUseCase.GetUserActiveInCategory(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err1 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err2 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err3 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToProfile(user, userbadges, catthreads, thread))
}

func (handler UserController) GetDashboardController(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetUsersController(ctx)
	threadreport, err1 := handler.ThreadReportUseCase.GetThreadReports(ctx)
	threadqty, err2 := handler.ThreadUseCase.GetThreadQuantity(ctx)
	userqty, err3 := handler.UserUseCase.GetUsersQuantity(ctx)
	postqty, err4 := handler.CommentUseCase.GetPostQuantity(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err1 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err1)
	}
	if err2 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err2)
	}
	if err3 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err3)
	}
	if err4 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err3)
	}
	fmt.Println("ThreadQty ", threadqty)
	fmt.Println("UserQty ", userqty)
	fmt.Println("PostQty ", postqty)
	fmt.Println("ThreadReport ", threadreport)
	return controllers.NewSuccessResponse(c, responses.ToDashboard(user, threadreport, userqty, threadqty, postqty))
}

func (handler UserController) GetSettingController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetUserSetting(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToUserSetting(user))
}
