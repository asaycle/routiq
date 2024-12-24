package usecase

import (
	"context"
	"time"

	"github.com/asaycle/motify.git/pkg/domain/entity"
	"github.com/asaycle/motify.git/pkg/domain/repository"
	"github.com/asaycle/motify.git/pkg/lib/filter"
	"golang.org/x/xerrors"
)

type TouringUsecase interface {
	CreateTouring(ctx context.Context, routeID string, title string, date time.Time, note string, score int, tags []*entity.Tag) (*entity.Touring, error)
	ListTourings(ctx context.Context, filter string) ([]*entity.Touring, error)
}

type TouringUsecaseImpl struct {
	touringRepo repository.TouringRepository
	routeRepo   repository.RouteRepository
	txManager   repository.TransactionManager
}

func NewTouringUsecaseImpl(touringRepo repository.TouringRepository, routeRepo repository.RouteRepository, txManager repository.TransactionManager) *TouringUsecaseImpl {
	return &TouringUsecaseImpl{
		touringRepo: touringRepo,
		routeRepo:   routeRepo,
		txManager:   txManager,
	}
}

func (u *TouringUsecaseImpl) CreateTouring(ctx context.Context, routeID string, title string, date time.Time, note string, score int, tags []*entity.Tag) (*entity.Touring, error) {
	touring := entity.NewTouring(routeID, title, date, note, score, tags)
	err := u.txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		if err := u.touringRepo.CreateTouring(ctx, tx, touring); err != nil {
			return xerrors.Errorf("Error CreateTouring: %w", err)
		}
		for _, tag := range touring.Tags {
			touringTag := entity.NewTouringTag(touring.ID, tag.ID)
			if err := u.touringRepo.CreateTouringTag(ctx, tx, touringTag); err != nil {
				return xerrors.Errorf("Error CreateTouringTag: %w", err)
			}
			routeTagCount := entity.NewRouteTagCount(touring.RouteID, tag.ID)
			if err := u.routeRepo.IncrementRouteTagCount(ctx, tx, routeTagCount); err != nil {
				return xerrors.Errorf("Error IncrementRouteTagCount: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("Failed Transaction: %w", err)
	}
	return touring, nil
}

func (u *TouringUsecaseImpl) ListTourings(ctx context.Context, filterStr string) ([]*entity.Touring, error) {
	touringFilter, err := filter.ParseListTouringsFilter(filterStr)
	if err != nil {
		return nil, xerrors.Errorf("Error ParseListTouringsFilter: %w", err)
	}

	var tourings []*entity.Touring
	err = u.txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		var err error
		tourings, err = u.touringRepo.ListTourings(ctx, tx, touringFilter)
		if err != nil {
			return xerrors.Errorf("Error ListTourings: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("Failed Transaction: %w", err)
	}
	return tourings, nil
}

var _ TouringUsecase = &TouringUsecaseImpl{}
