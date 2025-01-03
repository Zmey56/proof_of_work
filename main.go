package main

import (
	"fmt"

	"github.com/Zmey56/proof_of_work/blockchain"
)

func main() {
	// Create a new blockchain with a specified difficulty level
	bc := blockchain.NewBlockchain(4)

	fmt.Println("Adding blocks to the chain...")
	// Add new blocks to the blockchain
	bc.AddBlock("First block")
	bc.AddBlock("Second block")
	bc.AddBlock("Third block")

	// Display all blocks in the chain
	fmt.Println("Blockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\nPrevious Hash: %s\nData: %s\nHash: %s\nNonce: %d\n\n",
			block.Index, block.PreviousHash, block.Data, block.Hash, block.Nonce)
	}

	// Validate the blockchain integrity
	isValid := bc.IsValid()
	if isValid {
		fmt.Println("Blockchain is valid!")
	} else {
		fmt.Println("Blockchain is corrupted!")
	}
}
