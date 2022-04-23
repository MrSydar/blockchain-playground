package main

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"time"
)

type Chain struct {
	Blocks []Block
}

var chain *Chain = &Chain{
	Blocks: []Block{
		{
			PrevBlockHash: "",
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

	if err := rsa.VerifyPKCS1v15(payeePublicKey, 0, transactionJSON, signature); err != nil {
		return fmt.Errorf("failed to validate transaction signature: %v", err)
	}

	prevBlockHash, err := c.Blocks[len(c.Blocks)-1].GetHash()
	if err != nil {
		return fmt.Errorf("failed to get previous block hash: %v", err)
	}

	block := Block{
		PrevBlockHash: prevBlockHash,
		Transaction:   transaction,
		TimeSpamp:     time.Now().Unix(),
	}
	c.Blocks = append(c.Blocks, block)

	return nil
}

func (c *Chain) GetLastBlock() Block {
	return c.Blocks[len(c.Blocks)-1]
}
