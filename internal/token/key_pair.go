package token

import (
	"fmt"
	"os"
	"time"

	"github.com/anish-sinha1/sign-token-rsa/internal/util/errors"
	"github.com/anish-sinha1/sign-token-rsa/internal/util/helper"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
}

func CreateKeyPair(privateKeyBytes, publicKeyBytes []byte) *KeyPair {
	return &KeyPair{
		PrivateKey: privateKeyBytes,
		PublicKey:  publicKeyBytes,
	}
}

func (k *KeyPair) CreateSignedToken(ttl time.Duration) (string, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
	clientId := os.Getenv("CLIENT_ID")
	issuer := os.Getenv("ISSUER")
	if err != nil {
		errors.EnvironmentVariableError(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(k.PrivateKey)
	if err != nil {
		errors.ParseError(err)
	}
	nonce, err := helper.CreateNonce()
	if err != nil {
		errors.NonceError(err)
	}
	claims := make(jwt.MapClaims)
	claims["iss"] = issuer
	claims["sub"] = "example_user_sub"
	claims["aud"] = clientId
	claims["exp"] = time.Now().Add(ttl).Unix()
	claims["iat"] = time.Now().Unix()
	claims["nonce"] = nonce
	claims["acr"] = "0"
	claims["amr"] = "pwd"
	claims["azp"] = clientId
	claims["auth_time"] = time.Now().Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		errors.SignTokenError(err)
	}
	return token, nil
}

func (k *KeyPair) ValidateTokenSignature(token string) (interface{}, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(k.PublicKey)
	if err != nil {
		errors.ParseError(err)
	}
	res, err := jwt.Parse(token, func(signedJwt *jwt.Token) (interface{}, error) {
		if _, ok := signedJwt.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
		}
		return key, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	claims, ok := res.Claims.(jwt.MapClaims)
	if !ok || !res.Valid {
		errors.ValidateTokenError(err)
	}
	fmt.Println("Succeeded in validating token")
	return claims, nil
}
