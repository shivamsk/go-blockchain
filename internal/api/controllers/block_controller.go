package controllers

import (
	"go-blockchain/internal/constants"
	"go-blockchain/internal/dto"
	"go-blockchain/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BlockController struct {
	blockService services.IBlockService
	logger       *zap.SugaredLogger
}

func NewBlockController(blockService services.IBlockService, logger *zap.SugaredLogger) *BlockController {
	return &BlockController{
		blockService: blockService,
		logger:       logger,
	}
}

func (bc *BlockController) CreateBlock(c *gin.Context) {
	var content dto.BlockRequest

	if err := c.ShouldBindJSON(&content); err != nil {
		bc.logger.Errorw("Unable to bind request", "ErrorMessage", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	blockData, blockErr := bc.blockService.CreateBlock(c, &content)

	if blockErr != nil {
		c.Set(constants.ErrorMsg, blockErr.Error())
		c.Set(constants.StatusCode, http.StatusInternalServerError)
		return
	}

	c.Set(constants.ResponseValue, &blockData)
	c.Set(constants.StatusCode, http.StatusCreated)
	return
}

func (bc *BlockController) GetBlocks(c *gin.Context) {
	blocks, err := bc.blockService.GetBlocks(c)

	if err != nil {
		c.Set(constants.ErrorMsg, err.Error())
		c.Set(constants.StatusCode, http.StatusInternalServerError)
	}
	c.Set(constants.ResponseValue, &blocks)
	c.Set(constants.StatusCode, http.StatusCreated)
	return
}
