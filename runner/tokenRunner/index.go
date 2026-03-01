package tokenRunner

import (
	"crypto/rand"
	"encoding/hex"
)

// Generate returns a cryptographically secure random hex string using length bytes of randomness.
func Generate(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
