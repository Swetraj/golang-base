package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateResetToken() string {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
