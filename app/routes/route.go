package routes

import (
	"fgd-alterra-29/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.Use(middleware.Logger())
	jwtAuth := middleware.JWTWithConfig(cl.JwtConfig)

	e.POST("register", cl.UserController.RegisterController)
	e.POST("login", cl.UserController.LoginController)
	e.GET("users", cl.UserController.GetUserController, jwtAuth)
}
