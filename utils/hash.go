package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash generates a SHA-256 hash of the input data.
func Hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
