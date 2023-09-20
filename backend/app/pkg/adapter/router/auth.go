package router

import (
	"github.com/Bottlist/bottlist/pkg/adapter/handler"
	"github.com/labstack/echo/v4"
)

func NewAuthRouter(e *echo.Group, reqAuth *echo.Group, authHandler handler.AuthHandler) AuthRouter {
	auth := e.Group("/auth")

	return &authRouter{
		e:           auth,
		reqAuth:     reqAuth,
		authHandler: authHandler,
	}
}

type AuthRouter interface {
	Router()
}

type authRouter struct {
	e           *echo.Group
	reqAuth     *echo.Group
	authHandler handler.AuthHandler
}

func (h *authRouter) Router() {
	h.e.POST("/signup/user", h.authHandler.PostProvisionalSignup)
	h.e.GET("/signup/user", h.authHandler.GetProvisionalSignup)
	h.e.POST("/login/user", h.authHandler.PostProvisionalSignup)
}
