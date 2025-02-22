package router

import (
	"github.com/labstack/echo/v4"
	"phobyjun/controller"
)

func setAuthController(e *echo.Echo) {
	e.POST(APIAuthSignup, controller.SignUp)
	e.POST(APIAuth, controller.LogIn)
	e.DELETE(APIAuth, controller.LogOut)

	e.GET(APIAuth, controller.GetUserInfo)
}
