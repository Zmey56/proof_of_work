package blockchain

// Blockchain represents the chain of blocks.
type Blockchain struct {
	Blocks     []Block // List of blocks in the chain
	Difficulty int     // Difficulty level for mining
}

// NewBlockchain initializes a new blockchain with a genesis block.
func NewBlockchain(difficulty int) *Blockchain {
	genesisBlock := NewBlock(0, "", "Genesis Block", difficulty)
	return &Blockchain{
		Blocks:     []Block{genesisBlock},
		Difficulty: difficulty,
	}
}

// AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(len(bc.Blocks), previousBlock.Hash, data, bc.Difficulty)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// IsValid checks the integrity of the blockchain.
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		current := bc.Blocks[i]
		previous := bc.Blocks[i-1]

		// Validate hash linkage between blocks
		if current.PreviousHash != previous.Hash {
			return false
		}
	}
	return true
}
