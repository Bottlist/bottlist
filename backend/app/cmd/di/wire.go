//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Bottlist/bottlist/cmd/di/provider"
	"github.com/google/wire"
)

func InitializeApp() (*provider.App, error) {
	wire.Build(provider.DefaultAppSet)
	return nil, nil
}
