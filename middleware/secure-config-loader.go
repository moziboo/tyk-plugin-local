package middleware

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"plugin-dev/util/ctx"
	"plugin-dev/util/logger"
)

type ContextKey string

var config = map[string]string{}

func loadPrivateKeyFromEnv() (*rsa.PrivateKey, error) {
	privKeyPEM := os.Getenv("RSA_PRIVATE_KEY")
	block, _ := pem.Decode([]byte(privKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func decryptConfig(base64Message string) string {
	privateKey, err := loadPrivateKeyFromEnv()
	if err != nil {
		logger.Info("Error loading private key")
		return ""
	}

	// Decode the base64 message
	encryptedMessageBytes, err := base64.StdEncoding.DecodeString(base64Message)
	if err != nil {
		logger.Info("Error decoding base64 config value")
		return ""
	}

	// Decrypt the message using the private key
	decryptedValue, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedMessageBytes)
	if err != nil {
		logger.Info("Error decrypting config value")
		return ""
	}

	return string(decryptedValue)
}

func AddApiKeyToHeader(rw http.ResponseWriter, r *http.Request) {
	apikey, found := config["apiid-apikey"]

	if !found {
		definition := ctx.GetDefinition(r)
		decryptedKey := decryptConfig(definition.ConfigData["apikey"].(string))
		config["apiid-apikey"] = decryptedKey
		//r.Header.Set("ApiKey", decryptedKey)
		ctx := r.Context()
		ctxWithValue := context.WithValue(ctx, ContextKey("secrets"), decryptedKey)
		r2 := r.WithContext(ctxWithValue)
		*r = *r2
	} else {
		//r.Header.Set("ApiKey", apikey)
		ctx := r.Context()
		ctxWithValue := context.WithValue(ctx, ContextKey("secrets"), apikey)
		r2 := r.WithContext(ctxWithValue)
		*r = *r2
	}
}

func CheckContext(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	secret := ctx.Value(ContextKey("secrets"))
	logger.Info(secret.(string))
}
