package pow

import (
	"fmt"
	"strings"

	"github.com/Zmey56/proof_of_work/utils"
)

// ProofOfWork holds data for mining a block.
type ProofOfWork struct {
	Data       string // Data to be hashed
	Difficulty int    // Number of leading zeros required in the hash
}

// Mine performs the mining process to find a valid nonce.
func (pow *ProofOfWork) Mine() (string, int) {
	nonce := 0
	prefix := strings.Repeat("0", pow.Difficulty)

	for {
		// Create a string by combining data and the nonce
		input := fmt.Sprintf("%s:%d", pow.Data, nonce)
		hash := utils.Hash(input)

		// Log progress every 100,000 iterations
		if nonce > 0 && nonce%100000 == 0 {
			utils.LogInfo("Nonce: %d, Hash: %s", nonce, hash)
		}

		// Check if the hash satisfies the difficulty condition
		if strings.HasPrefix(hash, prefix) {
			return hash, nonce
		}

		nonce++
	}
}
