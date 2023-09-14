package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewHelloRouter(e *echo.Echo) HelloRouter {
	hello := e.Group("/hello")
	return &helloRouter{
		e: hello,
	}
}

type HelloRouter interface {
	Router()
}

type helloRouter struct {
	e *echo.Group
}

func (h *helloRouter) Router() {
	h.e.GET("", hello)
	h.e.GET("/two", hello2)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func hello2(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!2")
}
