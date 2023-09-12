package usecase

type AuthUsecase interface {
	CreateProvisionalUser(*CreateProvisionalUserInput) error
}

func NewAuthUsecase() AuthUsecase {
	return &authUsecase{}
}

type authUsecase struct {
}

type CreateProvisionalUserInput struct {
}

func (a *authUsecase) CreateProvisionalUser(input *CreateProvisionalUserInput) error {

	return nil
}
