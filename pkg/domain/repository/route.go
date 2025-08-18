package repository

import (
	"github.com/asaycle/routiq/pkg/domain/entity"
	geojson "github.com/paulmach/go.geojson"
	"golang.org/x/xerrors"
)

type Route struct {
	ID                string `db:"id"`
	Name              string `db:"name"`
	Description       string `db:"description"`
	FeatureCollection []byte `db:"feature"`
	UserID            string `db:"user_id"`
}

func fromRouteEntity(route *entity.Route) (*Route, error) {
	featureCollectionJSON, err := route.FeatureCollection.MarshalJSON()
	if err != nil {
		return nil, xerrors.Errorf("Error Feature MarshalJSON: %v", err)
	}
	return &Route{
		ID:                route.ID,
		Name:              route.DisplayName,
		Description:       route.Description,
		FeatureCollection: featureCollectionJSON,
		UserID:            route.UserID,
	}, nil
}

func toRouteEntity(model *Route) (*entity.Route, error) {
	featureCollection, err := geojson.UnmarshalFeatureCollection(model.FeatureCollection)
	if err != nil {
		return nil, xerrors.Errorf("Failed UnmarshalFeature: %v", err)
	}
	return &entity.Route{
		ID:                model.ID,
		DisplayName:       model.Name,
		Description:       model.Name,
		FeatureCollection: featureCollection,
		UserID:            model.UserID,
	}, nil
}
