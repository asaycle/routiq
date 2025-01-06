package usecase

import (
	"context"
	"log"
	"log/slog"

	"github.com/asaycle/routiq.git/pkg/domain/entity"
	"github.com/asaycle/routiq.git/pkg/domain/repository"
	"github.com/asaycle/routiq.git/pkg/lib/maps"
	geojson "github.com/paulmach/go.geojson"
	"golang.org/x/xerrors"
)

//go:generate mockgen -source=route_usecase.go -destination=mocks/route_usecase.go -package=mocks
type RouteUsecase interface {
	CreateRoute(ctx context.Context, name string, desc string, geoJSON string) (*entity.Route, error)
	GetRoute(ctx context.Context, name string) (*entity.Route, error)
	ListRoutes(ctx context.Context) ([]*entity.Route, error)
}

type RouteUsecaseImpl struct {
	repo      repository.RouteRepository
	txManager repository.TransactionManager
}

func NewRouteUsecaseImpl(repo repository.RouteRepository, txManager repository.TransactionManager) *RouteUsecaseImpl {
	return &RouteUsecaseImpl{repo: repo, txManager: txManager}
}

func (u *RouteUsecaseImpl) CreateRoute(ctx context.Context, name string, desc string, geoJSON string) (*entity.Route, error) {
	feature, err := geojson.UnmarshalFeature([]byte(geoJSON))
	if err != nil {
		return nil, xerrors.Errorf("Failed Unmarshal Feature Collection: %v", err)
	}
	route := entity.NewRoute(name, desc, feature)
	slog.Info("CreateRoute", slog.Any("route", route))
	err = u.txManager.RunWithReadWriteTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		if err := u.repo.CreateRoute(ctx, tx, route); err != nil {
			return xerrors.Errorf("Error CreateRoute: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("Failed ReadWriteTx: %w", err)
	}
	return route, nil
}

func (u *RouteUsecaseImpl) GetRoute(ctx context.Context, name string) (*entity.Route, error) {
	log.Printf("[usecase] GetRoute1: %v", name)
	var route *entity.Route
	var tagCounts []*entity.RouteTagCount
	err := u.txManager.RunWithReadOnlyTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		var err error
		route, err = u.repo.GetRoute(ctx, tx, name)
		if err != nil {
			log.Printf("[usecase] Error GetRoute: %v", err)
			return xerrors.Errorf("Error GetRoute: %v", err)
		}
		tagCounts, err = u.repo.GetRouteTagCounts(ctx, tx, route.ID)
		if err != nil {
			return xerrors.Errorf("Error GetRouteTagCounts: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("Failed ReadOnlyTx: %w", err)
	}
	url, err := maps.GenerateGoogleMapsURL(route.Feature)
	if err != nil {
		return nil, xerrors.Errorf("Error GenGoogleMapURL: %v", err)
	}
	route.MapURL = url
	route.TagCounts = tagCounts
	return route, nil
}

func (u *RouteUsecaseImpl) ListRoutes(ctx context.Context) ([]*entity.Route, error) {
	var routes []*entity.Route
	err := u.txManager.RunWithReadOnlyTx(ctx, func(ctx context.Context, tx repository.Tx) error {
		var err error
		routes, err = u.repo.ListRoutes(ctx, tx)
		if err != nil {
			return xerrors.Errorf("Error ListRoutes: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("Failed ReadOnlyTx: %w", err)
	}
	return routes, nil
}

var _ RouteUsecase = &RouteUsecaseImpl{}
