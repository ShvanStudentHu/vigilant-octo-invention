package main

import (
	"fmt"
	"encoding/json"
	// "github.com/ShvanStudentHu/vigilant-octo-invention/internal/key"
	// "github.com/ShvanStudentHu/vigilant-octo-invention/internal/crypto"
	"github.com/ShvanStudentHu/vigilant-octo-invention/vault"
	"os"
    "github.com/gin-gonic/gin"
	"github.com/ShvanStudentHu/vigilant-octo-invention/api"
	"github.com/joho/godotenv"

)

type Information struct {
	Name string
	Lastname string
	Address string
	Age int 
	PhoneNumber string
	Role string

}

type Key struct {
	ID string
	created_at string

}

type vaultServer struct {
	vaultClient: *api.Client
	keyName: string
}



func main() {
client, _ := vault.CreateVaultClient()
	
vault.SetToken(client)
	
r := gin.Default()

keyName := os.Getenv("ENCRYPT_KEY")

api.CreateKeyRoute(r , client, keyName)


r.Run(":8080")
// vault.CreateTransitKey(client, )


personInfo := Information{
	Name:        "Sarah",
    Lastname:    "smith",
    Address:     "somewhere in the world",
    Age:         25,
    PhoneNumber: "06529032",
}

plainData, err := json.Marshal(personInfo)
if err != nil {
	panic(err)
}

ciphertext, _ := vault.EncryptWithVaultKey(client, os.Getenv("ENCRYPT_KEY"), string(plainData))

fmt.Println("Encrypted personInfo:", ciphertext)


// newkey, err := vault.RotateTransitKey(client, os.Getenv("ENCRYPT_KEY"))
// if err != nil {
// 	fmt.Errorf("Rotation Failed")
// }
// fmt.Println(newkey)

	// fmt.Println("Plain JSON:", string(plainData))

	// fmt.Printf( "%+v/n",personInfo)

	// plaintext := []byte("hey you!")

	// key, err := key.GenerateKey(32)
	// if err != nil {
	// 	panic(err) // <placeholder>
	// }
	// fmt.Println("plaintext:", string(plaintext))

	// encryptedText, err := crypto.EncryptData(plainData, key)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// fmt.Println("the key is:", key)
	// 	fmt.Println("Encrypted text:", encryptedText)


	// 	d, _ := crypto.DecryptData(encryptedText, key)
	// 	fmt.Println("decrypt: ", string(d))


}

