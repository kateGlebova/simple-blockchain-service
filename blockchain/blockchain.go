package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Blockchain struct {
	LastBlockHash string
	Storage       map[string]string
}

func NewBlockchain() *Blockchain {
	blockchain := Blockchain{LastBlockHash: "0", Storage: make(map[string]string)}
	return &blockchain
}

func (b *Blockchain) AddNewBlock(data ...string) string {
	block := NewBlock(b.LastBlockHash, data)
	err := b.putIntoStorage(block)
	if err == nil {b.LastBlockHash = block.BlockHash}
	return block.BlockHash
}

func (b *Blockchain) putIntoStorage(block Block) error{
	jsonBlock, err := json.Marshal(block)
	if err != nil {return err}
	b.Storage[block.BlockHash] = string(jsonBlock)
	return nil
}

func (b Blockchain) getFromStorage(hash string) (Block, error) {
	var block Block
	jsonBlock, found := b.Storage[hash]
	if !found { return block, errors.New(fmt.Sprintf("block with hash '%s' does not exist", hash)) }
	err := json.Unmarshal([]byte(jsonBlock), &block)
	return block, err
}

func (b Blockchain) GetNLastBlocks(n int) ([]Block) {
	blocksArray := make([]Block, 0, n)
	previousHash := b.LastBlockHash
	for i := 0; i < n; i++ {
		if previousHash == "0" { break }
		lastBlock, _ := b.getFromStorage(previousHash)
		blocksArray = append(blocksArray, lastBlock)
		previousHash = lastBlock.PreviousBlockHash
	}
	return blocksArray
}