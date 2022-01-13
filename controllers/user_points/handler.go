package userpoints

import (
	userpoint "fgd-alterra-29/business/user_points"
	"fgd-alterra-29/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserPointController struct {
	UserPointUseCase userpoint.UseCase
}

func NewUserPointController(userpointUC userpoint.UseCase) *UserPointController {
	return &UserPointController{
		UserPointUseCase: userpointUC,
	}
}

func (handler UserPointController) AddThreadPointController(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := handler.UserPointUseCase.AddThreadPointController(ctx, 1)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, user)
}

func (handler UserPointController) AddPostPointController(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := handler.UserPointUseCase.AddPostPointController(ctx, 1)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NoDataSuccessResponse(c, user)
}
