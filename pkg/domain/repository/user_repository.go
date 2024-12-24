package repository

import (
	"context"
	"database/sql"

	"github.com/asaycle/routiq.git/pkg/domain/entity"
	"golang.org/x/xerrors"
)

type UserRepository interface {
	CreateUser(ctx context.Context, tx Tx, user *entity.User) error
	FindUserByGoogleSub(ctx context.Context, tx Tx, sub string) (*entity.User, error)
}

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, tx Tx, user *entity.User) error {
	_, err := tx.NamedExecContext(ctx, `
INSERT INTO users (
	id,
	google_sub
) VALUES (
	:id,
	:google_sub
)`, fromUserEntity(user))
	if err != nil {
		return xerrors.Errorf("Error NamedExecContext: %v", err)
	}
	return nil
}

func (r *UserRepositoryImpl) FindUserByGoogleSub(ctx context.Context, tx Tx, sub string) (*entity.User, error) {
	model := &User{}
	err := tx.GetContext(ctx, model, `
SELECT
	id,
	google_sub
FROM users
WHERE google_sub = $1
LIMIT 1`,
		sub)
	if err != nil {
		switch {
		case xerrors.Is(err, sql.ErrNoRows):
			return nil, NotfoundError
		default:
			return nil, xerrors.Errorf("Error GetContext: %v", err)
		}
	}
	return toUserEntity(model), nil
}

var _ UserRepository = &UserRepositoryImpl{}
