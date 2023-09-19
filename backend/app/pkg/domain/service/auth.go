package service

import (
	"database/sql"
	"github.com/Bottlist/bottlist/pkg/domain/model"
	"github.com/Bottlist/bottlist/pkg/domain/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/mail"
	"time"
)

type Authservice interface {
	ComparePassword(password, passWordCompare string) error
	EmailValidator(EmailAddress string) (string, error)
	CheckEmailRegister(email string) error
	CheckToken(token string) (*model.ProvisionalUser, error)
}

func NewAuthservice(authRepository repository.AuthRepository) Authservice {
	return &authService{
		authRepository: authRepository,
	}
}

type authService struct {
	authRepository repository.AuthRepository
}

func (a *authService) ComparePassword(password, passWordCompare string) error {
	if password != passWordCompare {
		return echo.NewHTTPError(http.StatusBadRequest, "パスワードが一致しません")
	}
	return nil
}

func (a *authService) EmailValidator(email string) (string, error) {
	result, b := mail.ParseAddress(email)
	if b != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, "メールアドレスの形式が間違っています")
	}
	return result.Address, nil
}

func (a *authService) CheckEmailRegister(email string) error {
	_, err := a.authRepository.GetUserByEmail(email)
	if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "既に登録されているメールアドレスです")
	}

	if err != sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}
func (a *authService) CheckToken(token string) (*model.ProvisionalUser, error) {
	user, err := a.authRepository.GetProvisionalUserByToken(token)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "不正なリクエストです")
	}
	now, err := getNowTime()
	if err != nil {
		return nil, err
	}
	if now.After(user.ExpiredAt) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "tokenが無効です")
	}
	return user, nil
}

func getNowTime() (time.Time, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	now := time.Now().In(jst)
	return now, nil
}
