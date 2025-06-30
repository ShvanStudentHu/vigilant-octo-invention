package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func EncryptData(data, key []byte) ([]byte, error) {

block, err := aes.NewCipher(key)
if err != nil {
	return nil, fmt.Errorf("failed to chop data in sizes of 16 bytes")
}

aesGCM, err := cipher.NewGCM(block)
if err != nil {
	return nil, fmt.Errorf("failed!!")
}

nonce := make([]byte, aesGCM.NonceSize())
rand.Read(nonce)

encryptedText := aesGCM.Seal(nil, nonce, data, nil)

result := append(nonce, encryptedText...)

return result, nil 

}

func DecryptData(data, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)

	aesGCM, _ := cipher.NewGCM(block)

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, _ := aesGCM.Open(nil, nonce, ciphertext, nil)

	return plaintext, nil
	
}