package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"go-blockchain/internal/dto"
	"strconv"
	"time"

	"go.uber.org/zap"
)

type (
	IBlockService interface {
		CreateBlock(ctx context.Context, content *dto.BlockRequest) (*dto.Block, error)
		GetBlocks(ctx context.Context) (*[]dto.Block, error)
	}

	BlockServiceImpl struct {
		logger     *zap.SugaredLogger
		blockchain []dto.Block
	}
)

func NewBlockServiceImpl(logger *zap.SugaredLogger) *BlockServiceImpl {
	return &BlockServiceImpl{
		logger:     logger,
		blockchain: make([]dto.Block, 0),
	}
}

func (bs *BlockServiceImpl) CreateBlock(ctx context.Context, content *dto.BlockRequest) (*dto.Block, error) {

	if len(bs.blockchain) == 0 {
		bs.createGenesisBlock()
		return &bs.blockchain[0], nil
	}
	prevBlock := bs.blockchain[len(bs.blockchain)-1]

	newBlock, err := bs.generateBlock(&prevBlock, content)

	if err != nil {
		return nil, err
	}

	if bs.isBlockValid(*newBlock, prevBlock) {
		bs.blockchain = append(bs.blockchain, *newBlock)
	} else {
		return nil, errors.New("Not Valid Block")
	}

	return newBlock, nil

}

func (bs *BlockServiceImpl) GetBlocks(ctx context.Context) (*[]dto.Block, error) {
	return &bs.blockchain, nil
}

func (bs *BlockServiceImpl) createGenesisBlock() {
	t := time.Now()
	genesisBlock := dto.Block{}
	genesisBlock = dto.Block{0, t.String(), nil, bs.calculateHash(&genesisBlock), ""}

	bs.blockchain = append(bs.blockchain, genesisBlock)
}

func (bs *BlockServiceImpl) generateBlock(oldBlock *dto.Block, content *dto.BlockRequest) (*dto.Block, error) {

	var newBlock dto.Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Content = content
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = bs.calculateHash(&newBlock)

	return &newBlock, nil
}

func (bs *BlockServiceImpl) calculateHash(block *dto.Block) string {

	blockData, err := json.Marshal(block.Content)
	bs.logger.Infof("BlockData %s", string(blockData))

	if err != nil {
		bs.logger.Error("Block Marshal Error " + err.Error())

	}
	record := strconv.Itoa(block.Index) + block.Timestamp + string(blockData) + block.PrevHash
	sha := sha256.New()
	sha.Write([]byte(record))

	// Appending Nil at the end
	hashed := sha.Sum(nil)
	encodedString := hex.EncodeToString(hashed)
	return encodedString
}

func (bs *BlockServiceImpl) isBlockValid(newBlock, oldBlock dto.Block) bool {

	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	newHash := bs.calculateHash(&newBlock)
	if newHash != newBlock.Hash {
		return false
	}

	return true
}
