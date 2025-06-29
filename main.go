package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

type Information struct {
	name string
	lastname string
	adress string
	age int 
	phoneNumber string
	role string

}

type Key struct {
	ID string
	created_at string


}


func newInfo(data Information) *Information {
	return &data
}

func GenerateKey() ([]byte, error){
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}

func encryptData(plaintext, key []byte) ([]byte, error) {

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

encryptedText := aesGCM.Seal(nil, nonce, plaintext, nil)


return encryptedText, nil


}

//data comes in
//




func main() {
	// personInfo := newInfo(Information{
	// 	name: "Sarah",
	// 	lastname: "smith",
	// 	adress: "somewhere in the world",
	// 	age: 25,
	// 	phoneNumber: "06529032",
	// })
	// // fmt.Printf( "%+v/n",personInfo)

	plaintext := []byte("hey you!")

	key, err := GenerateKey()
	if err != nil {
		panic(err) // <placeholder>
	}
	fmt.Println("plaintext:", string(plaintext))

	encryptedText, err := encryptData(plaintext, key)
		if err != nil {
			panic(err)
		}
		// fmt.Println("the key is:", key)
		fmt.Println("Encrypted text:", encryptedText)
}