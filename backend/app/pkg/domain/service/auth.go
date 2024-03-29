package service

import (
	"context"
	"database/sql"
	"github.com/Bottlist/bottlist/pkg/domain/model"
	"github.com/Bottlist/bottlist/pkg/domain/repository"
	"github.com/Bottlist/bottlist/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/mail"
	"strconv"
)

type IFAuthService interface {
	ComparePassword(password, passWordCompare string) error
	EmailValidator(EmailAddress string) (string, error)
	CheckEmailRegister(email string) error
	CheckToken(token string) (*model.ProvisionalUser, error)
	CreateCookie(ctx context.Context, userId int) (*http.Cookie, error)
	DeleteCookie(ctx context.Context, cookie *http.Cookie) (*http.Cookie, error)
}

func NewAuthService(authRepository *repository.AuthRepository, sessionRepository repository.SessionRepository) *AuthService {
	return &AuthService{
		authRepository:    authRepository,
		sessionRepository: sessionRepository,
	}
}

type AuthService struct {
	authRepository    repository.IFAuthRepository
	sessionRepository repository.SessionRepository
}

func (a *AuthService) ComparePassword(password, passWordCompare string) error {
	if password != passWordCompare {
		return echo.NewHTTPError(http.StatusBadRequest, "パスワードが一致しません")
	}
	return nil
}

func (a *AuthService) EmailValidator(email string) (string, error) {
	result, b := mail.ParseAddress(email)
	if b != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, "メールアドレスの形式が間違っています")
	}
	return result.Address, nil
}

func (a *AuthService) CheckEmailRegister(email string) error {
	_, err := a.authRepository.GetUserByEmail(email)
	if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "既に登録されているメールアドレスです")
	}

	if err != sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}
func (a *AuthService) CheckToken(token string) (*model.ProvisionalUser, error) {
	user, err := a.authRepository.GetProvisionalUserByToken(token)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "requestが不正です：", err)
	}
	now := utils.Now()
	if now.After(user.ExpiredAt) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "tokenが無効です")
	}
	return user, nil
}

func (a *AuthService) CreateCookie(ctx context.Context, userId int) (*http.Cookie, error) {
	value := utils.NewUUID()
	expires := utils.GetHourDuration(24)
	cookie := new(http.Cookie)
	cookie.Name = "session_id"
	cookie.Value = value
	cookie.Expires = utils.GetTimeDelay(expires)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.SameSite = http.SameSiteLaxMode
	err := a.sessionRepository.SetSession(ctx, value, strconv.Itoa(userId), expires)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return cookie, nil
}

func (a *AuthService) DeleteCookie(ctx context.Context, cookie *http.Cookie) (*http.Cookie, error) {
	err := a.sessionRepository.DeleteSession(ctx, cookie.Value)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	cookie.MaxAge = -1
	return cookie, nil
}
