package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	denys, err := NewWallet()
	if err != nil {
		panic(fmt.Errorf("failed to create denys wallet: %v", err))
	}

	bob, err := NewWallet()
	if err != nil {
		panic(fmt.Errorf("failed to create bob wallet: %v", err))
	}

	jaron, err := NewWallet()
	if err != nil {
		panic(fmt.Errorf("failed to create jaron wallet: %v", err))
	}

	if err = denys.SendMoney(10, bob.PublicKey); err != nil {
		panic(fmt.Errorf("failed to send money from denys to bob: %v", err))
	}

	if err = bob.SendMoney(50, denys.PublicKey); err != nil {
		panic(fmt.Errorf("failed to send money from bob to denys: %v", err))
	}

	if err = jaron.SendMoney(25, denys.PublicKey); err != nil {
		panic(fmt.Errorf("failed to send money from jaron to denys: %v", err))
	}

	chainJSON, err := json.MarshalIndent(*chain, "", "  ")
	if err != nil {
		panic(fmt.Errorf("failed to marshal chain: %v", err))
	}

	fmt.Printf("Chain: %v\n", string(chainJSON))
}
