package users

import (
	"fgd-alterra-29/business/users"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/users/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.UseCase
}

func NewUserController(userUseCase users.UseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (handler UserController) GetUserController(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetUserController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromListDomain(user))
}
