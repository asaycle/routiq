package handler

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	pb "github.com/asaycle/motify.git/api/proto/v1"
	"github.com/asaycle/motify.git/pkg/domain/repository"
	"github.com/asaycle/motify.git/pkg/domain/usecase"
	"github.com/asaycle/motify.git/pkg/infrastructure/db"
	"github.com/asaycle/motify.git/pkg/lib/session"
	"github.com/asaycle/motify.git/pkg/server/handlers"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	handlers.RegisterHandler(NewRouteHandler())
}

type RouteHandler struct {
	pb.UnimplementedRouteServiceServer

	useCase usecase.RouteUsecase
}

func NewRouteHandler() *RouteHandler {
	pgdb, err := db.NewPgDB("localhost", 5432, "root", "root", "motify")
	if err != nil {
		log.Panic("failed initialize pgdb", err)
	}
	repo := repository.NewRouteRepositoryImpl()
	txManager := repository.NewTransactionManager(pgdb)
	return &RouteHandler{
		useCase: usecase.NewRouteUsecaseImpl(repo, txManager),
	}
}

func (h *RouteHandler) Register(s *grpc.Server) {
	pb.RegisterRouteServiceServer(s, h)
}

func (h *RouteHandler) CreateRoute(ctx context.Context, req *pb.CreateRouteRequest) (*pb.Route, error) {
	fmt.Println(req.Route.Feature)
	route, err := h.useCase.CreateRoute(ctx, req.Route.DisplayName, req.Route.Description, req.Route.Feature)
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
