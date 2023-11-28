package handler

import "github.com/Bottlist/bottlist/pkg/usecase"

type IFUserHandler interface {
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

type UserHandler struct {
	userUsecase usecase.IFUserUsecase
}
