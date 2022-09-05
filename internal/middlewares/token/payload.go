package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ExpiredToken = errors.New("expired token")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(id uuid.UUID, duration time.Duration) (*Payload, error) {

	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenId,
		UserId:    id,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil

}
func (payload *Payload) ValidateToken() error {
	if time.Now().After(payload.ExpiredAt) {
		return ExpiredToken
	}
	return nil
}
