package repository

import (
	"context"

	"github.com/asaycle/routiq/pkg/domain/entity"
	"golang.org/x/xerrors"
)

type TagRepository interface {
	CreateTag(ctx context.Context, tx Tx, tag *entity.Tag) error
	GetByIDs(ctx context.Context, tx Tx, ids []string) ([]*entity.Tag, error)
	ListTags(ctx context.Context, tx Tx) ([]*entity.Tag, error)
}

type TagRepositoryImpl struct {
}

// NewTagRepositoryImpl は PostgreSQL リポジトリを作成
func NewTagRepositoryImpl() *TagRepositoryImpl {
	return &TagRepositoryImpl{}
}

// ListTags はタグの一覧を取得
func (r *TagRepositoryImpl) ListTags(ctx context.Context, tx Tx) ([]*entity.Tag, error) {
	var models []*Tag
	// TODO: add filter
	err := tx.SelectContext(ctx, &models, "SELECT id, name FROM tags")
	if err != nil {
		return nil, err
	}

	tags := make([]*entity.Tag, len(models))
	for i, model := range models {
		tags[i] = toTagEntity(model)
	}
	return tags, nil
}

func (r *TagRepositoryImpl) GetByIDs(ctx context.Context, tx Tx, ids []string) ([]*entity.Tag, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var models []*Tag
	err := tx.SelectContext(ctx, &models, "SELECT id, name FROM tags WHERE id IN (:ids)", ids)
	if err != nil {
		return nil, xerrors.Errorf("Failed BindNamed: %w", err)
	}
	if len(models) != len(ids) {
		return nil, xerrors.Errorf("Not found tags: %v", ids)
	}
	tags := make([]*entity.Tag, len(models))
	for i, model := range models {
		tags[i] = toTagEntity(model)
	}
	return tags, nil
}

// CreateTag はタグを新規作成
func (r *TagRepositoryImpl) CreateTag(ctx context.Context, tx Tx, tag *entity.Tag) error {
	_, err := tx.NamedExecContext(ctx, "INSERT INTO tags (id, name) VALUES (:id, :name)", fromTagEntity(tag))
	if err != nil {
		return xerrors.Errorf("Failed NamedExecContext: %w", err)
	}
	return nil
}

var _ TagRepository = &TagRepositoryImpl{}
