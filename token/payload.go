package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTokenExpired = errors.New("token expired")
	ErrInvalidToken = errors.New("invalid token")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("failed to generate token id")
	}
	now := time.Now()
	return &Payload{
		ID:        id,
		Username:  username,
		IssuedAt:  now,
		ExpiresAt: now.Add(duration),
	}, nil
}

func (p *Payload) IsValid() error {
	if time.Now().After(p.ExpiresAt) {
		return ErrTokenExpired
	}
	return nil
}
