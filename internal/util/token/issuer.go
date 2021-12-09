package token

import (
	"io/ioutil"

	"github.com/anish-sinha1/sign-token-rsa/models"
)

const (
	privateKeyPath = "../../config/rsa/private_key.pem"
	publicKeyPath  = "../../config/rsa/public_key.pem"
)

func LoadKeyPair(privateKey, publicKey []byte) *models.KeyPair {
	return &models.KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

func TokenValidator() (*models.KeyPair, error) {
	privateKey, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, err
	}
	publicKey, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	return LoadKeyPair(privateKey, publicKey), nil
}
