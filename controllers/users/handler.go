package users

import (
	"fgd-alterra-29/business/badges"
	"fgd-alterra-29/business/categories"
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/business/users"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/users/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase     users.UseCase
	ThreadUseCase   threads.UseCase
	BadgeUseCase    badges.UseCase
	CategoryUseCase categories.UseCase
}

func NewUserController(userUC users.UseCase, threadUC threads.UseCase, badgeUC badges.UseCase, categoryUC categories.UseCase) *UserController {
	return &UserController{
		UserUseCase:     userUC,
		ThreadUseCase:   threadUC,
		BadgeUseCase:    badgeUC,
		CategoryUseCase: categoryUC,
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
	point, _ := strconv.Atoi(c.QueryParam("point"))

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetProfileController(ctx, id)
	thread, err1 := handler.ThreadUseCase.GetProfileThreads(ctx, id)
	badges, err2 := handler.BadgeUseCase.GetBadgesByPointController(ctx, point)
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
	return controllers.NewSuccessResponse(c, responses.ToProfile(user, badges, catthreads, thread))
}
