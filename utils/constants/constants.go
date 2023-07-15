package constants

import (
	"crypto/rand"
	"log"
)

// Ideally the mongo URI, SECRET_KEY and AES_KEY would be in the secret manager in AWS or GCP
const (
	BASE_URL_PLD = "http://pld-container:3000"
	MONGODB_URI  = "mongodb://mongodb:27017"
	SECRET_KEY   = "test-crabi"
)

// This key should be static
var (
	AES_KEY = generateRandomAES256Key()
)

func generateRandomAES256Key() []byte {
	key := make([]byte, 32) // 32 bytes for AES-256
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	log.Print(key)
	return key
}
