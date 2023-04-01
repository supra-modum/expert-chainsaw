package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func generateRandomSecret(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(randomBytes), nil
}

func main() {
	secret, err := generateRandomSecret(32)
	if err != nil {
		log.Fatal("Error generating secret:", err)
	}

	envFile := ".env"
	envContent := fmt.Sprintf("JWT_SECRET=%s\n", secret)

	err = ioutil.WriteFile(envFile, []byte(envContent), os.FileMode(0600))
	if err != nil {
		log.Fatal("Error writing secret to .env file:", err)
	}

	fmt.Println("JWT Secret generated and stored in .env file.")
}
