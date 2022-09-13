package utils

import (
	"encoding/hex"
	"math/rand"
)

// RandomString godoc
// Creates a random hex string of length n.
func RandomString(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
