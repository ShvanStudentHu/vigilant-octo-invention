package key

import (
	"crypto/rand"
)

func GenerateKey(size int) ([]byte, error){
	key := make([]byte, size)
	_, err := rand.Read(key)
	return key, err
}