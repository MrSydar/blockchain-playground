package main

type Transaction struct {
	Amount         float64 `json:"amount"`
	PayerPublicKey []byte  `json:"payer_public_key"` //person who is paying the money
	PayeePublicKey []byte  `json:"payee_public_key"` //person who is receiving the money
}
