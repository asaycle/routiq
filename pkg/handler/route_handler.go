package handler

import (
	"context"
	"log"
	"log/slog"

	pb "github.com/asaycle/routiq.git/api/proto/v1"
	"github.com/asaycle/routiq.git/pkg/domain/repository"
	"github.com/asaycle/routiq.git/pkg/domain/usecase"
	"github.com/asaycle/routiq.git/pkg/infrastructure/db"
	"github.com/asaycle/routiq.git/pkg/lib/config"
	"github.com/asaycle/routiq.git/pkg/lib/session"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RouteHandler struct {
	pb.UnimplementedRouteServiceServer
	cfg     *config.Config
	useCase usecase.RouteUsecase
}

func ActivateRouteHandler(s *grpc.Server, cfg *config.Config) {
	pb.RegisterRouteServiceServer(s, NewRouteHandler(cfg))
}

func NewRouteHandler(cfg *config.Config) *RouteHandler {
	pgdb, err := db.NewPgDB("localhost", 5432, "root", "root", "routiq")
	if err != nil {
		log.Panic("failed initialize pgdb", err)
	}
	repo := repository.NewRouteRepositoryImpl()
	txManager := repository.NewTransactionManager(pgdb)
	return &RouteHandler{
		useCase: usecase.NewRouteUsecaseImpl(repo, txManager),
	}
}

func (h *RouteHandler) CreateRoute(ctx context.Context, req *pb.CreateRouteRequest) (*pb.Route, error) {
	route, err := h.useCase.CreateRoute(ctx, req.Route.DisplayName, req.Route.Description, req.Route.GeoJson)
	if err != nil {
		return nil, xerrors.Errorf("Error CreateRoute: %v", err)
	}
	pbRoute, err := route.ToProto()
	if err != nil {
		return nil, xerrors.Errorf("Error ToProto: %v", err)
	}
	return pbRoute, nil
}

func (h *RouteHandler) GetRoute(ctx context.Context, req *pb.GetRouteRequest) (*pb.Route, error) {
	log.Printf("[handler] GetRoute: %s", req.Id)
	route, err := h.useCase.GetRoute(ctx, req.Id)
	if err != nil {
		return nil, xerrors.Errorf("Error GetRoute: %v", err)
	}
	pbRoute, err := route.ToProto()
	if err != nil {
		return nil, xerrors.Errorf("Error ToProto: %v", err)
	}
	log.Printf("GetRoute: %v", pbRoute)
	return pbRoute, nil
}

func (h *RouteHandler) ListRoutes(ctx context.Context, req *pb.ListRoutesRequest) (*pb.ListRoutesResponse, error) {
	ss := session.GetSession(ctx)
	slog.Info("session", slog.Any("o", ss))
	if ss.Role == session.RoleGuest {
		return nil, status.Error(codes.PermissionDenied, "failed")
	}

	routes, err := h.useCase.ListRoutes(ctx)
	if err != nil {
		return nil, err
	}
	pbRoutes := make([]*pb.Route, len(routes))
	for i, route := range routes {
		pbRoute, err := route.ToProto()
		if err != nil {
			return nil, xerrors.Errorf("Error ToProto: %v", err)
		}
		pbRoutes[i] = pbRoute
	}
	return &pb.ListRoutesResponse{Routes: pbRoutes}, nil
}

var _ pb.RouteServiceServer = &RouteHandler{}
