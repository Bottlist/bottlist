package provider

import (
	"github.com/Bottlist/bottlist/pkg/adapter/handler"
	"github.com/google/wire"
)

var DefaultAppSet = wire.NewSet(
	DefaultUsecaseSet,
	DefaultHandlerSet,
	NewApp,
)

type App struct {
	AuthHandler handler.AuthHandler
}

func NewApp(
	AuthHandler handler.AuthHandler,
) *App {
	return &App{
		AuthHandler: AuthHandler,
	}
}
