package repository

import (
	"context"
	"log"

	"github.com/asaycle/motify.git/pkg/domain/entity"
	"golang.org/x/xerrors"
)

type RouteRepository interface {
	CreateRoute(ctx context.Context, tx Tx, route *entity.Route) error
	IncrementRouteTagCount(ctx context.Context, tx Tx, routeTagCount *entity.RouteTagCount) error
	GetRoute(ctx context.Context, tx Tx, name string) (*entity.Route, error)
	GetRouteTagCounts(ctx context.Context, tx Tx, routeID string) ([]*entity.RouteTagCount, error)
	ListRoutes(ctx context.Context, tx Tx) ([]*entity.Route, error)
}

type RouteRepositoryImpl struct {
}

// NewRouteRepositoryImpl は PostgreSQL リポジトリを作成
func NewRouteRepositoryImpl() *RouteRepositoryImpl {
	return &RouteRepositoryImpl{}
}

// CreateRoute はルートを新規作成
func (r *RouteRepositoryImpl) CreateRoute(ctx context.Context, tx Tx, route *entity.Route) error {
	modelRoute, err := fromRouteEntity(route)
	if err != nil {
		return xerrors.Errorf("Error fromRouteEntity: %v", err)
	}
	log.Print(modelRoute)
	_, err = tx.NamedExecContext(ctx, `
INSERT INTO routes (
	id,
	name,
	description,
	feature,
	user_id
) VALUES (
	:id,
	:name,
	:description,
	:feature,
	:user_id
)`, modelRoute)
	if err != nil {
		return xerrors.Errorf("Error NamedExecContext: %v", err)
	}
	return nil
}

func (r *RouteRepositoryImpl) IncrementRouteTagCount(ctx context.Context, tx Tx, routeTagCount *entity.RouteTagCount) error {
	model := fromRouteTagCountEntity(routeTagCount)
	_, err := tx.NamedExecContext(ctx, `
INSERT INTO route_tag_counts (
	id,
	route_id,
	tag_id,
	count
)
VALUES (
	:id,
	:route_id,
	:tag_id,
	1
)
ON CONFLICT (route_id, tag_id)
DO UPDATE SET count = route_tag_counts.count + 1;
`,
		model,
	)
	if err != nil {
		return xerrors.Errorf("Error NamedExecCountext: %w", err)
	}
	return nil
}

func (r *RouteRepositoryImpl) GetRoute(ctx context.Context, tx Tx, name string) (*entity.Route, error) {
	modelRoute := &Route{}
	err := tx.GetContext(ctx, modelRoute, `
SELECT
	id,
	name,
	description,
	feature
FROM routes
WHERE id = $1
LIMIT 1`,
		name)
	if err != nil {
		return nil, xerrors.Errorf("Error GetContext: %v", err)
	}
	route, err := toRouteEntity(modelRoute)
	if err != nil {
		return nil, xerrors.Errorf("Error toRouteEntity: %v", err)
	}
	return route, nil
}

// ListTags はタグの一覧を取得
func (r *RouteRepositoryImpl) ListRoutes(ctx context.Context, tx Tx) ([]*entity.Route, error) {
	var models []*Route
	err := tx.SelectContext(ctx, &models, `SELECT id, name, description, feature FROM routes`)
	if err != nil {
		return nil, err
	}

	routes := make([]*entity.Route, len(models))
	for i, model := range models {
		route, err := toRouteEntity(model)
		if err != nil {
			return nil, xerrors.Errorf("Error toRouteEntity: %w", err)
		}
		routes[i] = route
	}
	return routes, nil
}

func (r *RouteRepositoryImpl) GetRouteTagCounts(ctx context.Context, tx Tx, routeID string) ([]*entity.RouteTagCount, error) {
	var models []*RouteTagCount
	err := tx.SelectContext(ctx, &models, `
SELECT
	route_tag_counts.id AS id,
	tags.id AS tag_id,
	tags.name AS tag_name,
	route_tag_counts.count AS count
FROM route_tag_counts
INNER JOIN tags
ON route_tag_counts.tag_id = tags.id
WHERE 
	route_tag_counts.route_id = $1
`,
		routeID)
	if err != nil {
		return nil, err
	}

	counts := make([]*entity.RouteTagCount, len(models))
	for i, model := range models {
		counts[i] = toRouteTagCountEntity(model)
	}
	return counts, nil
}

var _ RouteRepository = &RouteRepositoryImpl{}
