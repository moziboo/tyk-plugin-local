package middleware

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"plugin-dev/util/ctx"
	"plugin-dev/util/logger"
)

func loadAESKeyFromEnv() ([]byte, error) {
	keyBase64 := os.Getenv("SECRETS_AES_KEY")
	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return nil, err
	}
	if len(key) != 32 {
		return nil, fmt.Errorf("AES key must be 32 bytes long")
	}
	return key, nil
}

func decryptAES(key []byte, cryptoText string) ([]byte, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

func decryptConfig(base64Message string) string {
	key, err := loadAESKeyFromEnv()
	if err != nil {
		logger.Info("Error loading AES key")
		return ""
	}

	// Decrypt the message using the AES key
	decryptedValue, err := decryptAES(key, base64Message)
	if err != nil {
		logger.Info("Error decrypting config value")
		return ""
	}

	return string(decryptedValue)
}

func AddApiKeyToHeader(rw http.ResponseWriter, r *http.Request) {
	definition := ctx.GetDefinition(r)
	decryptedKey := decryptConfig(definition.ConfigData["apikey"].(string))
	r.Header.Set("ApiKey", decryptedKey)
}
