package usecase

import (
	"context"
	"github.com/Bottlist/bottlist/pkg/domain/service"
)

type AuthUsecase interface {
	CreateProvisionalUser(ctx context.Context, input *CreateProvisionalUserInput) error
}

func NewAuthUsecase(authService service.Authservice) AuthUsecase {
	return &authUsecase{
		authService: authService,
	}
}

type authUsecase struct {
	authService service.Authservice
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

	return nil
}
