package provider

import (
	"github.com/Bottlist/bottlist/external/mail"
	"github.com/Bottlist/bottlist/external/mysql"
	"github.com/Bottlist/bottlist/external/redis"
	"github.com/Bottlist/bottlist/pkg/domain/repository"
	"github.com/google/wire"
)

var DefaultInfraSet = wire.NewSet(
	mysql.NewMySQLConnector,
	redis.NewRedisConnector,
	mail.NewMailClient,
	repository.NewSessionRepository,
	repository.NewAuthRepository,
	repository.NewUserRepository,
)
