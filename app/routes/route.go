package routes

import (
	"fgd-alterra-29/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.GET("users", cl.UserController.GetUserController)
}
