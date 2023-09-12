package handler

import (
	"github.com/Bottlist/bottlist/pkg/usecase"
	"github.com/labstack/echo/v4"
)

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func (a *AuthHandler) ProvisionalSignup(c echo.Context) error {
	return nil
}
