package handler

import (
	"github.com/Bottlist/bottlist/api/gen"
	"github.com/Bottlist/bottlist/pkg/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthHandler interface {
	PostProvisionalSignup(c echo.Context) error
	GetProvisionalSignup(c echo.Context) error
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
	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "requestが不正です：", err)
	}
	createProvisionalUserInput := &usecase.CreateProvisionalUserInput{
		Email:           req.User.Email,
		FirstName:       req.User.FirstName,
		LastName:        req.User.LastName,
		FirstNameHira:   req.User.FirstNameHuri,
		LastNameHira:    req.User.LastNameHuri,
		ScreenName:      req.User.ScreenName,
		Birthday:        req.User.Birthday,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	}
	if err := a.authUsecase.CreateProvisionalUser(ctx, createProvisionalUserInput); err != nil {
		return err
	}
	res := &gen.PostAuthProvisionalSignupResponse{}
	return c.JSON(http.StatusOK, res)
}

func (a *authHandler) GetProvisionalSignup(c echo.Context) error {
	ctx := c.Request().Context()

	var req gen.GetAuthSignupUserRequest
	token := c.QueryParam("token")
	req.Token = token
	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "requestが不正です：", err)
	}
	createUserInput := &usecase.CreateUserInput{
		req.Token,
	}
	if err := a.authUsecase.CreateUser(ctx, createUserInput); err != nil {
		return err
	}
	res := &gen.GetAuthSignupUserResponse{}
	return c.JSON(http.StatusOK, res)
}
