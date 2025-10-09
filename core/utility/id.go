package utility

import (
	"crypto/rand"
	"encoding/hex"
)

func NewID(length int) string {
	if length <= 0 {
		return ""
	}

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
