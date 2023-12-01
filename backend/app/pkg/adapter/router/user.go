package router

import (
	"github.com/Bottlist/bottlist/pkg/adapter/handler"
	"github.com/labstack/echo/v4"
)

func NewUserRouter(noAuth *echo.Group, reqAuth *echo.Group, userHandler handler.UserHandler) UserRouter {
	reqAuth = reqAuth.Group("/users")
	return &userRouter{
		e:           noAuth,
		eAuth:       reqAuth,
		userHandler: userHandler,
	}
}

type UserRouter interface {
	Router()
}

type userRouter struct {
	e           *echo.Group
	eAuth       *echo.Group
	userHandler handler.UserHandler
}

func (u *userRouter) Router() {
	u.eAuth.GET("", u.userHandler.GetUsers)
}
