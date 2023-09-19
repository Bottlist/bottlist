package usecase

import (
	"context"
	"fmt"
	"github.com/Bottlist/bottlist/external/mail"
	"github.com/Bottlist/bottlist/pkg/domain/model"
	"github.com/Bottlist/bottlist/pkg/domain/repository"
	"github.com/Bottlist/bottlist/pkg/domain/service"
	"github.com/Bottlist/bottlist/pkg/types"
	"github.com/Bottlist/bottlist/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AuthUsecase interface {
	CreateProvisionalUser(ctx context.Context, input *CreateProvisionalUserInput) error
	CreateUser(ctx context.Context, input *CreateUserInput) error
}

func NewAuthUsecase(authService service.Authservice, authRepository repository.AuthRepository, mailClient *mail.Client) AuthUsecase {
	return &authUsecase{
		authService:    authService,
		authRepository: authRepository,
		mailClient:     mailClient,
	}
}

type authUsecase struct {
	authService    service.Authservice
	authRepository repository.AuthRepository
	mailClient     *mail.Client
}

type CreateProvisionalUserInput struct {
	Email           string
	FirstName       string
	LastName        string
	FirstNameHira   string
	LastNameHira    string
	ScreenName      string
	Birthday        string
	Password        string
	PasswordConfirm string
}

func (a *authUsecase) CreateProvisionalUser(ctx context.Context, input *CreateProvisionalUserInput) error {
	email, err := a.authService.EmailValidator(input.Email)
	if err != nil {
		return err
	}

	err = a.authService.ComparePassword(input.Password, input.PasswordConfirm)
	if err != nil {
		return err
	}

	err = a.authService.CheckEmailRegister(email)
	if err != nil {
		return err
	}

	token := utils.NewUUID()
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	expiredAt := time.Now().In(jst).Add(1 * time.Hour)
	birth := types.Date{}
	err = birth.UnmarshalJSON(input.Birthday)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "誕生日の変換に失敗しました")
	}
	hashPassword, err := utils.PasswordEncrypt(input.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	provisionalUser := &model.CreateProvisionalUser{
		Email:         email,
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		FirstNameHira: input.FirstNameHira,
		LastNameHira:  input.LastNameHira,
		ScreenName:    input.ScreenName,
		Birthday:      birth,
		Password:      hashPassword,
		Token:         token,
		ExpiredAt:     expiredAt,
	}
	err = a.authRepository.InsertProvisionalUser(provisionalUser)
	if err != nil {
		return err
	}
	// TODO urlを正規のものにする
	body := registration_mail_body(token, expiredAt)
	mails := []string{email}
	err = a.mailClient.SendMail(mails, "title", body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return nil
}

func registration_mail_body(url string, validateTime time.Time) string {
	expiredAt := validateTime.Format("2006-01-02 15:04:05")
	return fmt.Sprintf(
		`
bottlist に仮登録いただきありがとうございます。

ユーザー登録手続きはまだ完了しておりません。
以下のURLにアクセスして、ユーザー登録の完了をお願いいたします。
%s

本URLは %s まで有効です。
有効期限経過後は再度メールアドレス登録から行ってください。


bottlist 運営チーム`,
		url,
		expiredAt,
	)
}

type CreateUserInput struct {
	Token string
}

func (a *authUsecase) CreateUser(ctx context.Context, input *CreateUserInput) error {
	provisionalUser, err := a.authService.CheckToken(input.Token)
	if err != nil {
		return err
	}
	user := &model.UserCreate{
		Email:         provisionalUser.Email,
		FirstName:     provisionalUser.FirstName,
		LastName:      provisionalUser.LastName,
		FirstNameHira: provisionalUser.FirstNameHira,
		LastNameHira:  provisionalUser.LastNameHira,
		ScreenName:    provisionalUser.ScreenName,
		Birthday:      provisionalUser.Birthday,
		Password:      provisionalUser.Password,
		Image:         "default",
	}
	err = a.authRepository.InsertUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err = a.authRepository.DeleteProvisionalUser(provisionalUser.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return nil
}
