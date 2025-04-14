package handler

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	pb "github.com/asaycle/routiq.git/api/proto/v1"
	"github.com/asaycle/routiq.git/pkg/domain/repository"
	"github.com/asaycle/routiq.git/pkg/domain/usecase"
	"github.com/asaycle/routiq.git/pkg/infrastructure/db"
	"github.com/asaycle/routiq.git/pkg/lib/config"
	"github.com/asaycle/routiq.git/pkg/lib/session"
	"golang.org/x/xerrors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
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
	if err := validateCreateRouteRequest(req); err != nil {
		fmt.Println(err)
		return nil, err
	}
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

func validateCreateRouteRequest(req *pb.CreateRouteRequest) error {
	if req.Route == nil {
		return status.Error(codes.InvalidArgument, "Route is required")
	}
	var fieldViolations []*errdetails.BadRequest_FieldViolation
	route := req.Route
	if route.DisplayName == "" {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{
			Field:       "display_name",
			Description: "display_name is required",
		})
	}
	if route.Description == "" {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{
			Field:       "description",
			Description: "description is required",
		})
	}
	if route.GeoJson == "" {
		fieldViolations = append(fieldViolations, &errdetails.BadRequest_FieldViolation{
			Field:       "geo_json",
			Description: "geo_json is required",
		})
	}

	if len(fieldViolations) > 0 {
		slog.Error("Failed to add details to error", slog.Any("violations", fieldViolations))
		st := status.New(codes.InvalidArgument, "Validation failed")
		br := &errdetails.BadRequest{FieldViolations: fieldViolations}
		stWithDetails, err := st.WithDetails(br)
		if err != nil {
			slog.Error("Failed to add details to error", slog.Any("err", err))
		}
		slog.Error("error", slog.Any("err", stWithDetails.Err()))
		return stWithDetails.Err()
	}

	return nil
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
