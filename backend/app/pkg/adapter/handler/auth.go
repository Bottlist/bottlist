package handler

import (
	"github.com/Bottlist/bottlist/api/gen"
	"github.com/Bottlist/bottlist/pkg/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IFAuthHandler interface {
	GetProvisionalSignup(c echo.Context) error
	PostProvisionalSignup(c echo.Context) error
	PostAuthLoginUser(c echo.Context) error
	PostLogout(c echo.Context) error
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

type AuthHandler struct {
	authUsecase usecase.IFAuthUsecase
}

func (a *AuthHandler) GetProvisionalSignup(c echo.Context) error {
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

func (a *AuthHandler) PostProvisionalSignup(c echo.Context) error {
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

func (a *AuthHandler) PostAuthLoginUser(c echo.Context) error {
	ctx := c.Request().Context()

	var req gen.PostAuthLoginUserRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "requestのBindに失敗しました：", err)
	}
	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "requestが不正です：", err)
	}

	input := &usecase.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}
	cookie, err := a.authUsecase.LoginUser(ctx, input)
	if err != nil {
		return err
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (a *AuthHandler) PostLogout(c echo.Context) error {
	ctx := c.Request().Context()
	cookie, err := c.Cookie("session_id")
	if err != nil {

		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	deleteCookie, err := a.authUsecase.Logout(ctx, cookie)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	c.SetCookie(deleteCookie)
	return c.NoContent(http.StatusOK)
}
