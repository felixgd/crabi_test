package cypher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/zenazn/pkcs7pad"
)

func Encrypt(plainText string, key []byte) (string, error) {
	plaintext := []byte(plainText)

	// Add PKCS7 padding to the plaintext
	paddedText := pkcs7pad.Pad(plaintext, aes.BlockSize)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(paddedText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedText)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
