package repository

import (
	"context"

	"github.com/asaycle/routiq/pkg/domain/entity"
	"golang.org/x/xerrors"
)

type LocationRepository interface {
	CreatePref(ctx context.Context, tx Tx, route *entity.Pref) error
	CreateCity(ctx context.Context, tx Tx, route *entity.City) error
}

type locationRepositoryImpl struct {
}

func NewLocationRepository() LocationRepository {
	return &locationRepositoryImpl{}
}

// CreatePref はPrefを新規作成
func (r *locationRepositoryImpl) CreatePref(ctx context.Context, tx Tx, pref *entity.Pref) error {
	model, err := fromPrefEntity(pref)
	if err != nil {
		return xerrors.Errorf("Error fromCityEntity: %v", err)
	}
	_, err = tx.NamedExecContext(ctx, `
INSERT INTO prefs (
	id,
	display_name
) VALUES (
	:id,
	:display_name
)`, model)
	if err != nil {
		return xerrors.Errorf("Error NamedExecContext: %v", err)
	}
	return nil
}

// CreateCity はCityを新規作成
func (r *locationRepositoryImpl) CreateCity(ctx context.Context, tx Tx, city *entity.City) error {
	model, err := fromCityEntity(city)
	if err != nil {
		return xerrors.Errorf("Error fromCityEntity: %v", err)
	}
	_, err = tx.NamedExecContext(ctx, `
INSERT INTO cities (
	id,
	pref_id,
	display_name
) VALUES (
	:id,
	:pref_id,
	:display_name
)`, model)
	if err != nil {
		return xerrors.Errorf("Error NamedExecContext: %v", err)
	}
	return nil
}
