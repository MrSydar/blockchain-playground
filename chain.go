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
	Blocks []Block
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

	h256 := sha256.New()
	h256.Write(transactionJSON)
	if err := rsa.VerifyPKCS1v15(payeePublicKey, crypto.SHA256, h256.Sum(nil), signature); err != nil {
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

func (c *Chain) mine(nonce int) int {
	target := []byte{0, 0}

	fmt.Println("🪨️ Mining...")

	for solution := 1; ; solution++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%d", nonce+solution)))
		if reflect.DeepEqual(hash[:4], target) {
			fmt.Printf("🎉 Block mined!\nSolution: %v", solution)
			return solution
		}
	}
}

func (c *Chain) GetLastBlock() Block {
	return c.Blocks[len(c.Blocks)-1]
}
