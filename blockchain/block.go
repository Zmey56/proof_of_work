package blockchain

import (
	"github.com/Zmey56/proof_of_work/pow"
)

// Block represents a single block in the blockchain.
type Block struct {
	Index        int    // Block index in the chain
	PreviousHash string // Hash of the previous block
	Data         string // Block data
	Hash         string // Current block hash
	Nonce        int    // Nonce value found during mining
}

// NewBlock creates a new block using the Proof-of-Work algorithm.
func NewBlock(index int, previousHash, data string, difficulty int) Block {
	powInstance := pow.ProofOfWork{
		Data:       data,
		Difficulty: difficulty,
	}

	// Perform mining to find a valid nonce and hash
	hash, nonce := powInstance.Mine()

	return Block{
		Index:        index,
		PreviousHash: previousHash,
		Data:         data,
		Hash:         hash,
		Nonce:        nonce,
	}
}
