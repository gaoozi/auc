package util

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

func GeneratePasswordHash(password string) ([]byte, error) {
  salt, err := generateRandBytes(16)
  if err != nil {
    return nil, err
  }

  key := argon2.IDKey([]byte(password), salt, 1, 64 * 1024, 2, 32)
  return key, nil
}

func VerifyPassword(password string, password_hash string) bool {
  key, _ := GeneratePasswordHash(password)
  if string(key) != password_hash {
    return false
  }
  return true
}

func generateRandBytes(n uint32) ([]byte, error) {
  bytes := make([]byte, n)
  _, err := rand.Read(bytes)
  if err != nil {
    return nil, err
  }

  return bytes, nil
}
