package repository

import (
	"log/slog"

	"github.com/asaycle/routiq.git/pkg/domain/entity"
	geojson "github.com/paulmach/go.geojson"
	"golang.org/x/xerrors"
)

type Route struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Feature     []byte `db:"feature"`
	UserID      string `db:"user_id"`
}

func fromRouteEntity(route *entity.Route) (*Route, error) {
	featureJSON, err := route.Feature.MarshalJSON()
	if err != nil {
		return nil, xerrors.Errorf("Error Feature MarshalJSON: %v", err)
	}
	return &Route{
		ID:          route.ID,
		Name:        route.Name,
		Description: route.Description,
		Feature:     featureJSON,
		UserID:      route.UserID,
	}, nil
}

func toRouteEntity(model *Route) (*entity.Route, error) {
	slog.Info("to route entity", slog.Any("model", model.Feature))
	feature, err := geojson.UnmarshalFeature(model.Feature)
	if err != nil {
		return nil, xerrors.Errorf("Failed UnmarshalFeature: %v", err)
	}
	return &entity.Route{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Name,
		Feature:     feature,
		UserID:      model.UserID,
	}, nil
}
