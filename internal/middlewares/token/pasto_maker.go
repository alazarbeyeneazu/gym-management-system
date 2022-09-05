package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

var InvalidToken = errors.New("invalid token")
var AuthroizationType = "bearer "

type PastoToken struct {
	pasto       *paseto.V2
	symtericKey []byte
}

func ErrorHelper(err string) error {
	return errors.New(err)
}

func NewPastoMaker(symerticKey string) (Maker, error) {
	if len(symerticKey) != chacha20poly1305.KeySize {
		return nil, ErrorHelper(fmt.Sprintf("Invalid key size %d the key must be %d ", len(symerticKey), chacha20poly1305.KeySize))
	}
	maker := &PastoToken{
		pasto:       paseto.NewV2(),
		symtericKey: []byte(symerticKey),
	}
	return maker, nil

}

func (pasto *PastoToken) CreateToken(userId uuid.UUID, duration time.Duration) (string, *Payload, error) {

	payload, err := NewPayload(userId, duration)
	if err != nil {
		return "", nil, err
	}

	token, err := pasto.pasto.Encrypt(pasto.symtericKey, payload, nil)
	if err != nil {
		return "", nil, err
	}
	return AuthroizationType + token, payload, nil
}
func (pasto *PastoToken) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := pasto.pasto.Decrypt(token, pasto.symtericKey, payload, nil)
	if err != nil {
		return nil, InvalidToken
	}
	err = payload.ValidateToken()
	if err != nil {
		return nil, err
	}
	return payload, nil

}
