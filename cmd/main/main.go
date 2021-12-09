package main

import (
	"fmt"
	"time"

	"github.com/anish-sinha1/sign-token-rsa/internal/util/token"
	"github.com/anish-sinha1/sign-token-rsa/models"
)

var (
	Validator *models.KeyPair
)

func init() {
	v, err := token.TokenValidator()
	if err != nil {
		panic(err)
	}
	Validator = v
}

func main() {
	token, err := Validator.CreateSignedToken(24 * time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
	claims, err := Validator.ValidateTokenSignature(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(claims)
}
