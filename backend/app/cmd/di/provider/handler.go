package provider

import (
	"github.com/Bottlist/bottlist/pkg/adapter/handler"
	"github.com/google/wire"
)

var DefaultHandlerSet = wire.NewSet(
	handler.NewAuthHandler,
	handler.NewUserHandler,
)
