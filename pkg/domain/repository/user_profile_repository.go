package repository

import (
	"context"

	"github.com/asaycle/routiq.git/pkg/domain/entity"
	"golang.org/x/xerrors"
)

type UserProfile struct {
	ID        string `db:"id"`
	UserID    string `db:"user_id"`
	AvatarURL string `db:"avatar_url"`
}

func fromUserProfileEntity(e *entity.UserProfile) *UserProfile {
	return &UserProfile{
		ID:        e.ID,
		UserID:    e.UserID,
		AvatarURL: e.AvatarURL,
	}
}

func toUserProfileEntity(m *UserProfile) *entity.UserProfile {
	return &entity.UserProfile{
		ID:        m.ID,
		UserID:    m.UserID,
		AvatarURL: m.AvatarURL,
	}
}

type UserProfileRepository interface {
	CreateUserProfile(ctx context.Context, tx Tx, user *entity.UserProfile) error
}

type UserProfileRepositoryImpl struct{}

func NewUserProfileRepositoryImpl() *UserProfileRepositoryImpl {
	return &UserProfileRepositoryImpl{}
}

func (r *UserProfileRepositoryImpl) CreateUserProfile(ctx context.Context, tx Tx, profile *entity.UserProfile) error {
	_, err := tx.NamedExecContext(ctx, `
INSERT INTO user_profiles (
	id,
	user_id,
	avatar_url
) VALUES (
	:id,
	:user_id,
	:avatar_url
)`, fromUserProfileEntity(profile))
	if err != nil {
		return xerrors.Errorf("Error NamedExecContext: %w", err)
	}
	return nil
}
