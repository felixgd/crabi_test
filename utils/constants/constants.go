package constants

import (
	"crypto/rand"
)

// Ideally the mongo URI, SECRET_KEY and AES_KEY would be in the secret manager in AWS or GCP
const (
	BASE_URL_PLD = "http://localhost:3000"
	MONGODB_URI  = "mongodb://localhost:27017"
	SECRET_KEY   = "test-crabi"
)

var (
	AES_KEY = generateRandomAES256Key()
)

func generateRandomAES256Key() []byte {
	key := make([]byte, 32) // 32 bytes for AES-256
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return key
}
