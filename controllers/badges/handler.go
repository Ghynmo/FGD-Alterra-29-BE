package badges

import (
	"fgd-alterra-29/business/badges"
	"fgd-alterra-29/controllers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BadgeController struct {
	BadgeUseCase badges.UseCase
}

func NewBadgeController(badgeUseCase badges.UseCase) *BadgeController {
	return &BadgeController{
		BadgeUseCase: badgeUseCase,
	}
}

func (handler BadgeController) GetBadgesByUserController(c echo.Context) error {
	point, _ := strconv.Atoi(c.Param("point"))
	ctx := c.Request().Context()

	badge, err := handler.BadgeUseCase.GetBadgesByUserController(ctx, point)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, badge)
}
