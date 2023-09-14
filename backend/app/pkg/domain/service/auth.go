package service

import "github.com/Bottlist/bottlist/pkg/domain/repository"

type Authservice interface {
}

func NewAuthservice(authRepository repository.AuthRepository) Authservice {
	return &authUsecase{
		authRepository: authRepository,
	}
}

type authUsecase struct {
	authRepository repository.AuthRepository
}
