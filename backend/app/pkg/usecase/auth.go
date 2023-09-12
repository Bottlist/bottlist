package usecase

type AuthUsecase interface {
}

type auth struct {
}

func NewAuthUsecase() AuthUsecase {
	return &auth{}
}
