package provider

import (
	"github.com/Bottlist/bottlist/middlewares"
	"github.com/Bottlist/bottlist/pkg/adapter/handler"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var DefaultAppSet = wire.NewSet(
	DefaultUsecaseSet,
	DefaultHandlerSet,
	DefaultInfraSet,
	DefaultServiceSet,
	DefaultMiddlewareSet,
	NewApp,
)

type App struct {
	AuthHandler *handler.AuthHandler
	UserHandler *handler.UserHandler
	Session     *middlewares.SessionMiddleware
	Db          *sqlx.DB
}

func NewApp(
	AuthHandler *handler.AuthHandler,
	UserHandler *handler.UserHandler,
	Session *middlewares.SessionMiddleware,
	Db *sqlx.DB,
) *App {
	return &App{
		AuthHandler: AuthHandler,
		UserHandler: UserHandler,
		Session:     Session,
		Db:          Db,
	}
}
