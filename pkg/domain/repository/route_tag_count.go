package repository

import "github.com/asaycle/routiq/pkg/domain/entity"

type RouteTagCount struct {
	ID      string `db:"id"`
	RouteID string `db:"route_id"`
	TagID   string `db:"tag_id"`
	TagName string `db:"tag_name"`
	Count   int    `db:"count"`
}

func fromRouteTagCountEntity(e *entity.RouteTagCount) *RouteTagCount {
	return &RouteTagCount{
		ID:      e.ID,
		RouteID: e.RouteID,
		TagID:   e.TagID,
		TagName: e.TagName,
		Count:   e.Count,
	}
}

func toRouteTagCountEntity(m *RouteTagCount) *entity.RouteTagCount {
	return &entity.RouteTagCount{
		ID:      m.ID,
		RouteID: m.RouteID,
		TagID:   m.TagID,
		TagName: m.TagName,
		Count:   m.Count,
	}
}
