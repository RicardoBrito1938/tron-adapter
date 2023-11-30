package main

import (
	"fmt"

	"github.com/blocktree/tron-adapter/tron"
)

func main() {
	// Initialize WalletManager
	wm := tron.NewWalletManager()

	// Generate a new address
	address, err := wm.GenerateAddress()
	if err != nil {
		fmt.Println("Error generating address:", err)
		return
	}

	// Print the generated address and private key
	fmt.Println("Address generated successfully!")
	fmt.Println("Address:", address["address"])
	fmt.Println("Private Key:", address["privateKey"])
}
