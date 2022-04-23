package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Block struct {
	Nonce int `json:"nonce"`

	PrevBlockHash []byte      `json:"prev_block_hash"`
	Transaction   Transaction `json:"transaction"`
	TimeSpamp     int64       `json:"time_stamp"`
}

func (b *Block) GetHash() ([]byte, error) {
	blockJSON, err := json.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal block: %v", err)
	}

	hash := sha256.Sum256(blockJSON)
	return hash[:], nil
}

func NewBlock(previousBlockHash []byte, transaction Transaction) *Block {
	return &Block{
		Nonce:         rand.Int(),
		PrevBlockHash: previousBlockHash,
		Transaction:   transaction,
		TimeSpamp:     time.Now().Unix(),
	}
}
