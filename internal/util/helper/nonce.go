package helper

import (
	"encoding/hex"

	"github.com/google/uuid"
)

func CreateNonce() (string, error) {
	nonce, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(nonce[:]), nil
}
