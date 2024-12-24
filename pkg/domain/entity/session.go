package entity

import (
	"time"

	"github.com/rs/xid"
)

type Session struct {
	ID           string
	UserID       string
	RefreshToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

func NewSession(userID string, refreshToken string, expiresAt time.Time) *Session {
	return &Session{
		ID:           xid.New().String(),
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt,
	}
}
