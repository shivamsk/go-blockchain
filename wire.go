//go:generate go run github.com/google/wire/cmd/wire
//go:build wireinject
// +build wireinject

package main

import (
	"go-blockchain/internal/api/controllers"
	"go-blockchain/internal/api/server"
	logger "go-blockchain/internal/common/log"

	"github.com/google/wire"
)

func InitApp() (App, error) {
	wire.Build(
		controllers.NewHealthController,
		logger.SugaredLogger,
		NewApp,
		server.NewServer,
		server.NewRouter,
	)

	return App{}, nil
}
