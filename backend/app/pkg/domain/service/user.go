package service

import "github.com/Bottlist/bottlist/pkg/domain/repository"

type IFUserService interface {
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

type UserService struct {
	userRepository repository.IFUserRepository
}
