package main

import (
	"fmt"
	"go-blockchain/internal/api/server"
	logger "go-blockchain/internal/common/log"
	"log"
)

type App struct {
	httpServer *server.Server
}

func NewApp(httpServer *server.Server) App {
	fmt.Println("NewApp")
	return App{
		httpServer: httpServer,
	}
}

func main() {

	fmt.Println("Hello hello")
	logger.InitLogger()
	app, err := InitApp()

	if err != nil {
		log.Panicf(fmt.Sprintf("Could not build app : #{err}"))
	}

	fmt.Println("Start hello")

	app.httpServer.Serve()

	fmt.Println("Main End")

}
