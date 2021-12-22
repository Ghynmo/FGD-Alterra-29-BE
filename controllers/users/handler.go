package users

import (
	"fgd-alterra-29/business/threads"
	userbadges "fgd-alterra-29/business/user_badges"
	"fgd-alterra-29/business/users"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/users/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase      users.UseCase
	ThreadUseCase    threads.UseCase
	UserBadgeUseCase userbadges.UseCase
}

func NewUserController(userUseCase users.UseCase, threadUseCase threads.UseCase, userbadgeUseCase userbadges.UseCase) *UserController {
	return &UserController{
		UserUseCase:      userUseCase,
		ThreadUseCase:    threadUseCase,
		UserBadgeUseCase: userbadgeUseCase,
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

	user, catthreads, err := handler.UserUseCase.GetProfileController(ctx, id)
	thread, err1 := handler.ThreadUseCase.GetProfileThreads(ctx, id)
	userbadges, err2 := handler.UserBadgeUseCase.GetUserBadge(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err1 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	if err2 != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToProfile(user, userbadges, catthreads, thread))
}
