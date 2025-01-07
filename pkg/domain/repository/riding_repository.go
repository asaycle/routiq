package repository

import (
	"context"
	"log/slog"

	"github.com/asaycle/routiq.git/pkg/domain/entity"
	pkgfilter "github.com/asaycle/routiq.git/pkg/lib/filter"
	"golang.org/x/xerrors"
)

type RidingRepository interface {
	CreateRiding(ctx context.Context, tx Tx, riding *entity.Riding) error
	CreateRidingTag(ctx context.Context, tx Tx, ridingTag *entity.RidingTag) error
	ListRidings(ctx context.Context, tx Tx, filter *pkgfilter.RidingFilter) ([]*entity.Riding, error)
}

type RidingRepositoryImpl struct{}

func NewRidingRepositoryImpl() *RidingRepositoryImpl {
	return &RidingRepositoryImpl{}
}

func (a *RidingRepositoryImpl) CreateRiding(ctx context.Context, tx Tx, riding *entity.Riding) error {
	model, err := fromRidingEntity(riding)
	if err != nil {
		return xerrors.Errorf("Error fromRidingEntity: %w", err)
	}
	_, err = tx.NamedExecContext(ctx, `
INSERT INTO ridings (
	id,
	route_id,
	title,
	date,
	score,
	note
) VALUES (
	:id,
	:route_id,
	:title,
	:date,
	:score,
	:note
)
	`, model)
	if err != nil {
		return xerrors.Errorf("Error NamedExecContext: %w", err)
	}
	return nil
}

func (a *RidingRepositoryImpl) CreateRidingTag(ctx context.Context, tx Tx, ridingTag *entity.RidingTag) error {
	model, err := fromRidingTagEntity(ridingTag)
	if err != nil {
		return xerrors.Errorf("Error fromRidingEntity: %w", err)
	}
	_, err = tx.NamedExecContext(ctx, `
INSERT INTO riding_tags (
	id,
	riding_id,
	tag_id
) VALUES (
	:id,
	:riding_id,
	:tag_id
)
	`, model)
	if err != nil {
		return xerrors.Errorf("Error NamedExecContext: %w", err)
	}
	return nil
}

func (a *RidingRepositoryImpl) ListRidings(ctx context.Context, tx Tx, filter *pkgfilter.RidingFilter) ([]*entity.Riding, error) {
	query := "SELECT * FROM ridings WHERE true"
	args := []interface{}{}
	if filter.RouteID != "" {
		query += " AND route_id = $1"
		args = append(args, filter.RouteID)
	}
	var models []*Riding
	slog.Info("QUERY", slog.Any("query", query), slog.Any("args", args))
	if err := tx.SelectContext(ctx, &models, query, args...); err != nil {
		return nil, xerrors.Errorf("Error SelectContext: %w", err)
	}

	ridings := make([]*entity.Riding, len(models))
	for i, model := range models {
		riding, err := toRidingEntity(model)
		if err != nil {
			return nil, xerrors.Errorf("Error toRidingEntity: %w", err)
		}
		ridings[i] = riding
	}
	return ridings, nil
}

var _ RidingRepository = &RidingRepositoryImpl{}
