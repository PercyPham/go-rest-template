package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

var secret = "secret"

// HashSHA256 returns hashed version of a string
func HashSHA256(s string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(s))
	hashed := hex.EncodeToString(h.Sum(nil))
	return hashed
}
