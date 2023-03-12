package randgenerator

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

// MustGenerateToString generates crypto random string encoded by base64.RawURLEncoding (without padding).
func MustGenerateToString(byteLength uint) string {
	randBytes := make([]byte, byteLength)
	_, err := io.ReadFull(rand.Reader, randBytes)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(randBytes)
}
