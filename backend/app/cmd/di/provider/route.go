package provider

import (
	"github.com/Bottlist/bottlist/pkg/adapter/router"
	"github.com/google/wire"
)

var DefaultRouteSet = wire.NewSet(
	router.NewAuthRouter,
)
