package provider

import (
	"github.com/Bottlist/bottlist/external/mysql"
	"github.com/Bottlist/bottlist/external/redis"
	"github.com/Bottlist/bottlist/pkg/infra"
	"github.com/google/wire"
)

var DefaultInfraSet = wire.NewSet(
	mysql.NewMySQLConnector,
	redis.NewRedisConnector,
	infra.NewSessionRepository,
)