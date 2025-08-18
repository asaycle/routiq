package repository

import (
	"context"
	"time"

	"github.com/asaycle/routiq/pkg/domain/entity"
	"golang.org/x/xerrors"
)

type Session struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	RefreshToken string    `db:"refresh_token"`
	ExpiresAt    time.Time `db:"expires_at"`
	CreatedAt    time.Time `db:"created_at"`
}

func fromSessionEntity(e *entity.Session) *Session {
	return &Session{
		ID:           e.ID,
		UserID:       e.UserID,
		RefreshToken: e.RefreshToken,
		ExpiresAt:    e.ExpiresAt,
		CreatedAt:    e.CreatedAt,
	}
}

type SessionRepository interface {
	CreateSession(ctx context.Context, tx Tx, session *entity.Session) error
}

type SessionRepositoryImpl struct{}

func NewSessionRepositoryImpl() *SessionRepositoryImpl {
	return &SessionRepositoryImpl{}
}

func (a *SessionRepositoryImpl) CreateSession(ctx context.Context, tx Tx, session *entity.Session) error {
	_, err := tx.NamedExecContext(ctx, `
INSERT INTO sessions (
	id,
	user_id,
	refresh_token,
	expires_at,
	created_at
) VALUES (
	:id,
	:user_id,
	:refresh_token,
	:expires_at,
	:created_at
)
	`, fromSessionEntity(session))
	if err != nil {
		return xerrors.Errorf("Error NamedExecContext: %w", err)
	}
	return nil
}
