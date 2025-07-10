package vault

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
)


func CreateVaultClient() (*api.Client, error) {
	
	config:= api.DefaultConfig()
	config.Address = "http://127.0.0.1:8200/"
	
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	
	return client, nil
}

func SetToken(client *api.Client) (error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load .env file")
	}
	
	token := os.Getenv("VAULT_TOKEN") 
	if token == ""{
		return fmt.Errorf("failed to set token")
	}
	
	client.SetToken(token)
	
	return nil
}

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load .env file")
	}
}


func CreateTransitKey(client *api.Client, keyName string) error {

	path := fmt.Sprintf("transit/keys/%s", keyName)
	

	data := map[string]interface{}{
		"type": "aes256-gcm96", // Or "rsa-2048", "ecdsa-p256", etc.
	}

	_, err := client.Logical().Write(path, data)
	if err != nil {
		return fmt.Errorf("failed to create transit key: %w", err)
	}

	return nil
}

func EncryptWithVaultKey(client *api.Client, keyName string, plaintext string) (string, error) {
	plaintextB64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

	data := map[string]interface{}{
		"plaintext": plaintextB64,
	}

	
	secret, err := client.Logical().Write(fmt.Sprintf("transit/encrypt/%s", keyName), data)
	if err != nil {
		return "", fmt.Errorf("encryption failed: %w", err)
	}

	
	ciphertext, ok := secret.Data["ciphertext"].(string)
	if !ok {
		return "", fmt.Errorf("ciphertext missing or invalid format")
	}

	return ciphertext, nil
}

func DecryptWithVaultKey(client *api.Client, keyName string, ciphertext string) (string, error) {
	data := map[string]interface{}{
		"ciphertext": ciphertext,
	}

	secret, err := client.Logical().Write(fmt.Sprintf("transit/decrypt/%s", keyName), data)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}


	plaintextB64, ok := secret.Data["plaintext"].(string)
	if !ok {
		return "", fmt.Errorf("plaintext missing or invalid format")
	}

	
	plaintextBytes, err := base64.StdEncoding.DecodeString(plaintextB64)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 plaintext: %w", err)
	}

	return string(plaintextBytes), nil
}
f