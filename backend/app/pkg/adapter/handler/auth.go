package handler

import (
	"github.com/Bottlist/bottlist/pkg/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler interface {
	PostProvisionalSignup(c echo.Context) error
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) AuthHandler {
	return &authHandler{authUsecase: authUsecase}
}

type authHandler struct {
	authUsecase usecase.AuthUsecase
}

type PostProvisionalSignupRequest struct {
}

func (a *authHandler) PostProvisionalSignup(c echo.Context) error {
	req := PostProvisionalSignupRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return nil
}
