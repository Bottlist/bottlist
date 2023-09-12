package router

import (
	"github.com/Bottlist/bottlist/pkg/adapter/handler"
	"github.com/labstack/echo/v4"
)

func NewAuthRouter(e *echo.Echo, authHandler *handler.AuthHandler) AuthRouter {
	auth := e.Group("/auth")

	return &authRouter{
		e:           auth,
		authHandler: authHandler,
	}
}

type AuthRouter interface {
	Router()
}

type authRouter struct {
	e           *echo.Group
	authHandler *handler.AuthHandler
}

func (h *authRouter) Router() {
	h.e.POST("/signup", h.authHandler.ProvisionalSignup)
}
