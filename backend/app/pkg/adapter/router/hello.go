package router

import (
	"github.com/Bottlist/bottlist/external/mail"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func NewHelloRouter(e *echo.Group) HelloRouter {
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
	client := mail.NewMailClient()
	mails := make([]string, 0)
	mails = append(mails, "kickryou2000@gmail.com")
	str := `first line

second line
third line`
	err := client.SendMail(mails, "title", str)
	log.Println(err)
	return c.String(http.StatusOK, "Hello, World!")
}

func hello2(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!2")
}
