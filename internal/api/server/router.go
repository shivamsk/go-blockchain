package server

import (
	"fmt"
	"go-blockchain/internal/api/controllers"
	logger "go-blockchain/internal/common/log"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RouterConfig struct {
	router           *gin.Engine
	healthController *controllers.HealthController
}

func NewRouter(
	healthController *controllers.HealthController,
) *RouterConfig {

	fmt.Println("NewRouter")
	return &RouterConfig{
		healthController: healthController,
	}
}

func (r *RouterConfig) addRoutes() *gin.Engine {
	router := gin.New()
	log := logger.NonSugaredLogger()

	log.Info("Starting Router",
		zap.String("ServiceName", "go-blockchain"))
	router.Use(
		controllers.HttpResponseInjector(),
		ginzap.GinzapWithConfig(log, &ginzap.Config{
			TimeFormat: time.RFC3339,
			UTC:        true,
			SkipPaths: []string{
				"/go-blockchain/health",
			},
		}),
	)

	health := controllers.NewHealthController(logger.SugarLogger)

	router.GET("/go-blockchain/health", health.Status)

	return router
}
