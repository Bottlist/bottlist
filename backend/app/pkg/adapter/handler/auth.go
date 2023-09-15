package handler

import (
	"github.com/Bottlist/bottlist/api/gen"
	"github.com/Bottlist/bottlist/pkg/usecase"
	"github.com/labstack/echo/v4"
	"log"
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

func (a *authHandler) PostProvisionalSignup(c echo.Context) error {
	ctx := c.Request().Context()

	var req gen.PostAuthProvisionalSignupRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "requestのBindに失敗しました：", err)
	}
	log.Println(req)
	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "requestが不正です：", err)
	}
	createProvisionalUserInput := &usecase.CreateProvisionalUserInput{
		Email:           req.Email,
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		FirstNameHira:   req.FirstNameHuri,
		LastNameHira:    req.LastNameHuri,
		ScreenName:      req.ScreenName,
		Birthday:        req.Birthday,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	}
	if err := a.authUsecase.CreateProvisionalUser(ctx, createProvisionalUserInput); err != nil {
		return err
	}
	res := &gen.PostAuthProvisionalSignupResponse{}
	return c.JSON(http.StatusOK, res)
}
