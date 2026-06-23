package models

import (
	"crypto/rand"
	"math/big"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYS0123456789"

func randomID(length int) string {
	b := make([]byte, length)

	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		b[i] = chars[n.Int64()]
	}

	return string(b)
}
