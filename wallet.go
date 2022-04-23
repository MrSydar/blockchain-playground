package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"fmt"
)

type Wallet struct {
	PublicKey  []byte `json:"public_key"`
	PrivateKey []byte `json:"private_key"`
}

func (w *Wallet) SendMoney(amount float64, payeePublicKey []byte) error {
	transaction := Transaction{
		Amount:         amount,
		PayerPublicKey: w.PublicKey,
		PayeePublicKey: payeePublicKey,
	}

	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction: %v", err)
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(w.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	h256 := sha256.Sum256(transactionJSON)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, h256[:])
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %v", err)
	}

	publicKey, err := x509.ParsePKCS1PublicKey(w.PublicKey)
	if err != nil {
		return fmt.Errorf("failed to parse public key: %v", err)
	}

	if err = chain.AddBlock(transaction, publicKey, signature); err != nil {
		return fmt.Errorf("failed to add block: %v", err)
	}

	return nil
}

func NewWallet() (*Wallet, error) {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	publickey := &privatekey.PublicKey

	return &Wallet{
		PrivateKey: x509.MarshalPKCS1PrivateKey(privatekey),
		PublicKey:  x509.MarshalPKCS1PublicKey(publickey),
	}, nil
}
