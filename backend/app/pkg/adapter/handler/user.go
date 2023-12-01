package handler

import (
	"net/http"

	"github.com/Bottlist/bottlist/pkg/usecase"
	"github.com/labstack/echo/v4"
)

type IFUserHandler interface {
	GetUsers(c echo.Context) error
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

type UserHandler struct {
	userUsecase usecase.IFUserUsecase
}

func (u *UserHandler) GetUsers(c echo.Context) error {
	//ctx := c.Request().Context()

	return c.NoContent(http.StatusOK)
}
