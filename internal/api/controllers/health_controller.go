package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HealthController struct {
	logger *zap.SugaredLogger
}

func NewHealthController(logger *zap.SugaredLogger) *HealthController {
	fmt.Println("NewHealthController")
	return &HealthController{logger: logger}
}

func (h *HealthController) Status(c *gin.Context) {
	version, err := os.ReadFile("deployedVersion.txt")

	if err != nil {
		h.logger.Error("Failed to read deployedVersion")
	}

	h.logger.Infof("Version %s", version)

	c.JSON(http.StatusOK, gin.H{"Version": string(version)})
}
