package userbadges

import (
	userbadges "fgd-alterra-29/business/user_badges"
	"fgd-alterra-29/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserBadgeController struct {
	UserBadgeUseCase userbadges.UseCase
}

func NewUserBadgeController(threadUseCase userbadges.UseCase) *UserBadgeController {
	return &UserBadgeController{
		UserBadgeUseCase: threadUseCase,
	}
}

func (handler UserBadgeController) GetProfileUserBadges(c echo.Context) error {
	ctx := c.Request().Context()
	id := 2

	userbadges, err := handler.UserBadgeUseCase.GetUserBadge(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, userbadges)
}
