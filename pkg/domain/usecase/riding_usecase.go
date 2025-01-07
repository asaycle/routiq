package usecase

import (
	"context"
	"time"

	"github.com/asaycle/routiq.git/pkg/domain/entity"
	"github.com/asaycle/routiq.git/pkg/domain/repository"
	"github.com/asaycle/routiq.git/pkg/lib/filter"
	"golang.org/x/xerrors"
)

type RidingUsecase interface {
	CreateRiding(ctx context.Context, routeID string, title string, date time.Time, note string, score int, tags []*entity.Tag) (*entity.Riding, error)
	ListRidings(ctx context.Context, filter string) ([]*entity.Riding, error)
}

type RidingUsecaseImpl struct {
	ridingRepo repository.RidingRepository
	routeRepo  repository.RouteRepository
	txManager  repository.TransactionManager
}

func NewRidingUsecaseImpl(ridingRepo repository.RidingRepository, routeRepo repository.RouteRepository, txManager repository.TransactionManager) *RidingUsecaseImpl {
	return &RidingUsecaseImpl{
		ridingRepo: ridingRepo,
		routeRepo:  routeRepo,
		txManager:  txManager,
	}
}

func (u *RidingUsecaseImpl) CreateRiding(ctx context.Context, routeID string, title string, date time.Time, note string, score int, tags []*entity.Tag) (*entity.Riding, error) {
	riding := entity.NewRiding(routeID, title, date, note, score, tags)
	err := u.txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		if err := u.ridingRepo.CreateRiding(ctx, tx, riding); err != nil {
			return xerrors.Errorf("Error CreateRiding: %w", err)
		}
		for _, tag := range riding.Tags {
			ridingTag := entity.NewRidingTag(riding.ID, tag.ID)
			if err := u.ridingRepo.CreateRidingTag(ctx, tx, ridingTag); err != nil {
				return xerrors.Errorf("Error CreateRidingTag: %w", err)
			}
			routeTagCount := entity.NewRouteTagCount(riding.RouteID, tag.ID)
			if err := u.routeRepo.IncrementRouteTagCount(ctx, tx, routeTagCount); err != nil {
				return xerrors.Errorf("Error IncrementRouteTagCount: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("Failed Transaction: %w", err)
	}
	return riding, nil
}

func (u *RidingUsecaseImpl) ListRidings(ctx context.Context, filterStr string) ([]*entity.Riding, error) {
	ridingFilter, err := filter.ParseListRidingsFilter(filterStr)
	if err != nil {
		return nil, xerrors.Errorf("Error ParseListRidingsFilter: %w", err)
	}

	var ridings []*entity.Riding
	err = u.txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		var err error
		ridings, err = u.ridingRepo.ListRidings(ctx, tx, ridingFilter)
		if err != nil {
			return xerrors.Errorf("Error ListRidings: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("Failed Transaction: %w", err)
	}
	return ridings, nil
}

var _ RidingUsecase = &RidingUsecaseImpl{}
