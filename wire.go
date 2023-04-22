//go:generate go run github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package main

import (
	"go-blockchain/internal/api/controllers"
	"go-blockchain/internal/api/server"
	logger "go-blockchain/internal/common/log"
	"go-blockchain/internal/services"

	"github.com/google/wire"
)

func InitApp() (App, error) {
	wire.Build(
		logger.SugaredLogger,
		controllers.NewHealthController,
		controllers.NewBlockController,
		services.NewBlockServiceImpl,
		NewApp,
		server.NewServer,
		server.NewRouter,
		wire.Bind(new(services.IBlockService), new(*services.BlockServiceImpl)),
	)

	return App{}, nil
}
