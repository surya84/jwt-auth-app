package initializers

import (
	"crypto/rsa"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

var PublicKey *rsa.PublicKey
var PrivateKey *rsa.PrivateKey

func LoadKeys() {
	PrivateKeyPath := os.Getenv("PRIVATE_KEY")
	publicKeyPath := os.Getenv("PUBLIC_KEY")
	privateKeyBytes, err := os.ReadFile(PrivateKeyPath)
	if err != nil {
		log.Fatal("Error reading private key:", err)
	}
	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		log.Fatal("Error parsing private key:", err)
	}

	// Load the public key
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal("Error reading public key:", err)
	}
	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		log.Fatal("Error parsing public key:", err)
	}

}
