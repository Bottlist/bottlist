package provider

import (
	"github.com/Bottlist/bottlist/pkg/domain/service"
	"github.com/google/wire"
)

var DefaultServiceSet = wire.NewSet(
	service.NewAuthservice,
)
