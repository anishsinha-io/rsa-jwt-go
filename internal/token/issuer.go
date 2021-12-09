package token

import (
	"io/ioutil"
)

const (
	privateKeyPath = "../../config/rsa/private_key.pem"
	publicKeyPath  = "../../config/rsa/public_key.pem"
)

func LoadKeyPair(privateKey, publicKey []byte) *KeyPair {
	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

func TokenValidator() (*KeyPair, error) {
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
