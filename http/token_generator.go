package http

import (
	"crypto/rand"
	"encoding/base64"
)

// RandToken generates a random @l length token.
func RandToken(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
