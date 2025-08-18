package usecase

import (
	"context"
	"time"

	"github.com/asaycle/routiq/pkg/domain/entity"
	"github.com/asaycle/routiq/pkg/domain/repository"
	"github.com/asaycle/routiq/pkg/lib/filter"
	"golang.org/x/xerrors"
)

type TouringUsecase interface {
	CreateTouring(ctx context.Context, routeID string, title string, date time.Time, note string, score int, tagIDs []string) (*entity.Touring, error)
	ListTourings(ctx context.Context, filter string) ([]*entity.Touring, error)
}

type TouringUsecaseImpl struct {
	touringRepo repository.TouringRepository
	routeRepo   repository.RouteRepository
	tagRepo     repository.TagRepository
	txManager   repository.TransactionManager
}

func NewTouringUsecaseImpl(touringRepo repository.TouringRepository, routeRepo repository.RouteRepository, txManager repository.TransactionManager) *TouringUsecaseImpl {
	return &TouringUsecaseImpl{
		touringRepo: touringRepo,
		routeRepo:   routeRepo,
		txManager:   txManager,
	}
}

func (u *TouringUsecaseImpl) CreateTouring(ctx context.Context, routeID string, title string, date time.Time, note string, score int, tagIDs []string) (*entity.Touring, error) {
	touring := entity.NewTouring(routeID, title, date, note, score, nil)
	err := u.txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		_, err := u.routeRepo.GetRoute(ctx, tx, routeID)
		if err != nil {
			return xerrors.Errorf("Error GetRoute: %w", err)
		}
		tags, err := u.tagRepo.GetByIDs(ctx, tx, tagIDs)
		if err != nil {
			return xerrors.Errorf("Error GetByIDs: %w", err)
		}
		if err := u.touringRepo.CreateTouring(ctx, tx, touring); err != nil {
			return xerrors.Errorf("Error CreateTouring: %w", err)
		}
		touring.Tags = tags
		for _, tag := range tags {
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
