package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Chain struct {
	Blocks []Block `json:"blocks"`
}

var chain *Chain = &Chain{
	Blocks: []Block{
		{
			PrevBlockHash: nil,
			TimeSpamp:     time.Now().Unix(),
			Transaction: Transaction{
				Amount:         21.0,
				PayerPublicKey: []byte("denys"),
				PayeePublicKey: []byte("pablo"),
			},
		},
	},
}

func (c *Chain) AddBlock(transaction Transaction, payeePublicKey *rsa.PublicKey, signature []byte) error {
	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction: %v", err)
	}

	h256 := sha256.Sum256(transactionJSON)
	if err := rsa.VerifyPKCS1v15(payeePublicKey, crypto.SHA256, h256[:], signature); err != nil {
		return fmt.Errorf("failed to validate transaction signature: %v", err)
	}

	prevBlockHash, err := c.Blocks[len(c.Blocks)-1].GetHash()
	if err != nil {
		return fmt.Errorf("failed to get previous block hash: %v", err)
	}

	block := NewBlock(prevBlockHash, transaction)
	c.mine(block.Nonce)
	c.Blocks = append(c.Blocks, *block)

	return nil
}

var miningTarget = []byte{0, 0}

func (c *Chain) mine(nonce int) int {
	if len(miningTarget) > 128 {
		panic(fmt.Errorf("mining target is too long"))
	}

	fmt.Println("🪨️ Mining...")

	for solution := 1; ; solution++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%d", nonce+solution)))
		if reflect.DeepEqual(hash[:len(miningTarget)], miningTarget) {
			fmt.Printf("🎉 Block mined!\nSolution: %v\n", solution)
			return solution
		}
	}
}

func (c *Chain) GetLastBlock() Block {
	return c.Blocks[len(c.Blocks)-1]
}
