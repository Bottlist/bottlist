package provider

import (
	"github.com/Bottlist/bottlist/pkg/usecase"
	"github.com/google/wire"
)

var DefaultUsecaseSet = wire.NewSet(
	usecase.NewAuthUsecase,
)
