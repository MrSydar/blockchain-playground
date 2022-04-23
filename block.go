package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Block struct {
	PrevBlockHash string      `json:"prev_block_hash"`
	Transaction   Transaction `json:"transaction"`
	TimeSpamp     int64       `json:"time_stamp"`
}

// function GetHash return a SHA256 hash of a marshalled Block struct
func (b *Block) GetHash() (string, error) {
	blockJSON, err := json.Marshal(b)
	if err != nil {
		return "", fmt.Errorf("failed to marshal block: %v", err)
	}

	hash := sha256.Sum256(blockJSON)
	return string(hash[:]), nil
}
