package usecase

import (
	"context"

	"github.com/asaycle/routiq.git/pkg/domain/entity"
	"github.com/asaycle/routiq.git/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type TagUsecase interface {
	CreateTag(ctx context.Context, name string) (*entity.Tag, error)
	ListTags(ctx context.Context) ([]*entity.Tag, error)
}

type TagUsecaseImpl struct {
	repo      repository.TagRepository
	txManager repository.TransactionManager
}

func NewTagUsecaseImpl(repo repository.TagRepository, txManager repository.TransactionManager) *TagUsecaseImpl {
	return &TagUsecaseImpl{repo: repo, txManager: txManager}
}

func (u *TagUsecaseImpl) CreateTag(ctx context.Context, name string) (*entity.Tag, error) {
	tag := entity.NewTag(name)
	err := u.txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		if err := u.repo.CreateTag(ctx, tx, tag); err != nil {
			return xerrors.Errorf("Error CreateTag: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("Failed ReadWriteTx: %w", err)
	}
	return tag, nil
}

func (u *TagUsecaseImpl) ListTags(ctx context.Context) ([]*entity.Tag, error) {
	var tags []*entity.Tag
	err := u.txManager.RunWithReadOnlyTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		var err error
		tags, err = u.repo.ListTags(ctx, tx)
		if err != nil {
			return xerrors.Errorf("Error ListTags: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("Failed ReadOnlyTx: %w", err)
	}
	return tags, nil
}

var _ TagUsecase = &TagUsecaseImpl{}
