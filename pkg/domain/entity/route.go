package entity

import (
	"log/slog"

	pb "github.com/asaycle/routiq.git/api/proto/v1"
	geojson "github.com/paulmach/go.geojson"
	"github.com/rs/xid"
	"golang.org/x/xerrors"
)

type Route struct {
	ID                string
	Name              string
	Description       string
	FeatureCollection *geojson.FeatureCollection
	MapURL            string
	UserID            string
	TagCounts         []*RouteTagCount
}

func (r *Route) ToProto() (*pb.Route, error) {
	var featureJSON []byte
	if r.FeatureCollection != nil {
		var err error
		featureJSON, err = r.FeatureCollection.MarshalJSON()
		if err != nil {
			return nil, xerrors.Errorf("Error MarshalJSON: %v", err)
		}
	}
	tagCounts := make([]*pb.TagCount, len(r.TagCounts))
	for i, tagCount := range r.TagCounts {
		slog.Info("tagCount", slog.Any("tag", tagCount))
		tagCounts[i] = tagCount.ToProto()
	}
	return &pb.Route{
		Id:           r.ID,
		DisplayName:  r.Name,
		Description:  r.Description,
		GeoJson:      string(featureJSON),
		GoogleMapUrl: r.MapURL,
		TagCounts:    tagCounts,
	}, nil
}

func NewRoute(name string, desc string, featureCollection *geojson.FeatureCollection) *Route {
	return &Route{
		ID:                xid.New().String(),
		Name:              name,
		Description:       desc,
		FeatureCollection: featureCollection,
		UserID:            "admin",
	}
}

type RouteTagCount struct {
	ID      string
	RouteID string
	TagID   string
	TagName string
	Count   int
}

func NewRouteTagCount(routeID string, tagID string) *RouteTagCount {
	return &RouteTagCount{
		ID:      xid.New().String(),
		RouteID: routeID,
		TagID:   tagID,
		Count:   0,
	}
}

func (r *RouteTagCount) ToProto() *pb.TagCount {
	return &pb.TagCount{
		Tag: &pb.Tag{
			Id:   r.TagID,
			Name: r.TagName,
		},
		Count: int32(r.Count),
	}
}
