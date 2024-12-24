package repository

import (
	"context"
	"log/slog"

	"github.com/asaycle/motify.git/pkg/domain/entity"
	pkgfilter "github.com/asaycle/motify.git/pkg/lib/filter"
	"golang.org/x/xerrors"
)

type TouringRepository interface {
	CreateTouring(ctx context.Context, tx Tx, touring *entity.Touring) error
	CreateTouringTag(ctx context.Context, tx Tx, touringTag *entity.TouringTag) error
	ListTourings(ctx context.Context, tx Tx, filter *pkgfilter.TouringFilter) ([]*entity.Touring, error)
}

type TouringRepositoryImpl struct{}

func NewTouringRepositoryImpl() *TouringRepositoryImpl {
	return &TouringRepositoryImpl{}
}

func (a *TouringRepositoryImpl) CreateTouring(ctx context.Context, tx Tx, touring *entity.Touring) error {
	model, err := fromTouringEntity(touring)
	if err != nil {
		return xerrors.Errorf("Error fromTouringEntity: %w", err)
	}
	_, err = tx.NamedExecContext(ctx, `
INSERT INTO tourings (
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

func (a *TouringRepositoryImpl) CreateTouringTag(ctx context.Context, tx Tx, touringTag *entity.TouringTag) error {
	model, err := fromTouringTagEntity(touringTag)
	if err != nil {
		return xerrors.Errorf("Error fromTouringEntity: %w", err)
	}
	_, err = tx.NamedExecContext(ctx, `
INSERT INTO touring_tags (
	id,
	touring_id,
	tag_id
) VALUES (
	:id,
	:touring_id,
	:tag_id
)
	`, model)
	if err != nil {
		return xerrors.Errorf("Error NamedExecContext: %w", err)
	}
	return nil
}

func (a *TouringRepositoryImpl) ListTourings(ctx context.Context, tx Tx, filter *pkgfilter.TouringFilter) ([]*entity.Touring, error) {
	query := "SELECT * FROM tourings WHERE true"
	args := []interface{}{}
	if filter.RouteID != "" {
		query += " AND route_id = $1"
		args = append(args, filter.RouteID)
	}
	var models []*Touring
	slog.Info("QUERY", slog.Any("query", query), slog.Any("args", args))
	if err := tx.SelectContext(ctx, &models, query, args...); err != nil {
		return nil, xerrors.Errorf("Error SelectContext: %w", err)
	}

	tourings := make([]*entity.Touring, len(models))
	for i, model := range models {
		touring, err := toTouringEntity(model)
		if err != nil {
			return nil, xerrors.Errorf("Error toTouringEntity: %w", err)
		}
		tourings[i] = touring
	}
	return tourings, nil
}

var _ TouringRepository = &TouringRepositoryImpl{}
