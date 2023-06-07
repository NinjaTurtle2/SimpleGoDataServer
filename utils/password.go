package utils

//Utility to hash password with random salt

import (
	"crypto/sha256"
	"crypto/rand"
	"fmt"
	"io"
)

func GenerateRandomSalt() string {
	randomBytes := make([]byte, 12)
	rand.Read(randomBytes)
	return fmt.Sprintf("%x", randomBytes)
}
// HashPassword hashes password with random salt
func HashPassword(password string, salt string) string {
	//Hash password with salt
	hash := sha256.New()
	io.WriteString(hash, password+salt)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
