package provider

import (
	"github.com/Bottlist/bottlist/middlewares"
	"github.com/google/wire"
)

var DefaultMiddlewareSet = wire.NewSet(
	middlewares.NewSessionMiddleware,
)
