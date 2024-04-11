package global

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptSHA256(s string) string {
	// Hash the password using SHA-256
	hasher := sha256.New()
	hasher.Write([]byte(s))
	hashedPassword := hasher.Sum(nil)

	// Convert the hashed password to hexadecimal
	encodedPassword := hex.EncodeToString(hashedPassword)

	return encodedPassword
}
