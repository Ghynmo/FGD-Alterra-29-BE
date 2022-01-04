package users

import (
	"fgd-alterra-29/business/users"
	"fgd-alterra-29/controllers"
	"fgd-alterra-29/controllers/users/request"
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

func (handler UserController) RegisterController(c echo.Context) error {
	var NewRegister request.Register
	c.Bind(&NewRegister)

	domain := NewRegister.FromRegister()

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.RegisterController(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToLoginResponse(user))
}

func (handler UserController) LoginController(c echo.Context) error {
	var NewLogin request.Login
	c.Bind(&NewLogin)

	domain := NewLogin.FromLogin()

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.LoginController(ctx, domain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.ToLoginResponse(user))
}

func (handler UserController) GetUserController(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := handler.UserUseCase.GetUserController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, responses.FromListDomain(user))
}
