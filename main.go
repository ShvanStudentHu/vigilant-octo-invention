package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ShvanStudentHu/vigilant-octo-invention/api"
	"github.com/ShvanStudentHu/vigilant-octo-invention/middleware"
	"github.com/ShvanStudentHu/vigilant-octo-invention/vault"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Information struct {
	Name string
	Lastname string
	Address string
	Age int 
	PhoneNumber string
	Role string

}

func main() {
	logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})
	client, _ := vault.CreateVaultClient()
	
	vault.SetToken(client)
	
	r := gin.Default()

 	r.Use(middleware.Logger(logger))
	keyName := os.Getenv("ENCRYPT_KEY")

	api.CreateKeyRoute(r , client, keyName)


	r.Run(":8080")



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
}