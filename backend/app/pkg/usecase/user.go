package usecase

import (
	"github.com/Bottlist/bottlist/pkg/domain/repository"
	"github.com/Bottlist/bottlist/pkg/domain/service"
)

type IFUserUsecase interface {
}

func NewUserUsecase(userService *service.UserService, userRepository *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userService:    userService,
		userRepository: userRepository,
	}
}

type UserUsecase struct {
	userService    service.IFUserService
	userRepository repository.IFUserRepository
}
