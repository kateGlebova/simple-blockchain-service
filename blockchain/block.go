package blockchain

import (
	"time"
	"crypto/sha256"
	"bytes"
	"fmt"
)

type Block struct {
	PreviousBlockHash string `json:"previous_block_hash"`
	Rows              []string `json:"rows"`
	Timestamp         time.Time `json:"timestamp"`
	BlockHash         string `json:"block_hash"`
}

func NewBlock(previousBlockHash string, rows []string) Block {
	timestamp := time.Now()
	blockHash := hashFields(previousBlockHash, rows, timestamp)
	block := Block{previousBlockHash, rows, timestamp, blockHash}
	return block
}

func hashFields(previousBlockHash string, rows []string, timestamp time.Time) string {
	var fields bytes.Buffer
	fields.Write([]byte(previousBlockHash))
	for _, row := range rows {
		fields.Write([]byte(row))
	}
	bytesTime, _ := timestamp.MarshalText()
	fields.Write(bytesTime)

	hash := sha256.Sum256(fields.Bytes())
	return fmt.Sprintf("%x", hash)
}