package handler

import (
	"context"
	"log"
	"time"

	pb "github.com/asaycle/routiq.git/api/proto/v1"
	"github.com/asaycle/routiq.git/pkg/domain/entity"
	"github.com/asaycle/routiq.git/pkg/domain/repository"
	"github.com/asaycle/routiq.git/pkg/domain/usecase"
	"github.com/asaycle/routiq.git/pkg/infrastructure/db"
	"github.com/asaycle/routiq.git/pkg/lib/config"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

func ActivateRidingHandler(s *grpc.Server, cfg *config.Config) {
	pb.RegisterRidingServiceServer(s, NewRidingHandler(cfg))
}

type RidingHandler struct {
	pb.UnimplementedRidingServiceServer
	cfg     *config.Config
	useCase usecase.RidingUsecase
}

func NewRidingHandler(cfg *config.Config) *RidingHandler {
	pgdb, err := db.NewPgDB("localhost", 5432, "root", "root", "routiq")
	if err != nil {
		log.Panic("failed initialize pgdb", err)
	}
	txManager := repository.NewTransactionManager(pgdb)
	return &RidingHandler{
		cfg: cfg,
		useCase: usecase.NewRidingUsecaseImpl(
			repository.NewRidingRepositoryImpl(),
			repository.NewRouteRepositoryImpl(),
			txManager,
		),
	}
}

func (h *RidingHandler) CreateRiding(ctx context.Context, req *pb.CreateRidingRequest) (*pb.Riding, error) {
	date := time.Date(int(req.Riding.Date.Year), time.Month(req.Riding.Date.Month), int(req.Riding.Date.Day), 0, 0, 0, 0, time.Local)
	tags := make([]*entity.Tag, len(req.Riding.Tags))
	for i, tag := range req.Riding.Tags {
		tags[i] = &entity.Tag{
			ID:   tag.Id,
			Name: tag.Name,
		}
	}
	riding, err := h.useCase.CreateRiding(ctx, req.Riding.RouteId, req.Riding.Title, date, req.Riding.Note, int(req.Riding.Score), tags)
	if err != nil {
		return nil, xerrors.Errorf("Error CreateRiding: %w", err)
	}
	pbRiding, err := riding.ToProto()
	if err != nil {
		return nil, xerrors.Errorf("Error ToProto: %w", err)
	}
	return pbRiding, nil
}

func (h *RidingHandler) ListRidings(ctx context.Context, req *pb.ListRidingsRequest) (*pb.ListRidingsResponse, error) {
	ridings, err := h.useCase.ListRidings(ctx, req.Filter)
	if err != nil {
		return nil, xerrors.Errorf("Error ListRidings: %w", err)
	}
	pbRidings := make([]*pb.Riding, len(ridings))
	for i, e := range ridings {
		pbRiding, err := e.ToProto()
		if err != nil {
			return nil, xerrors.Errorf("Error ToProto: %w", err)
		}
		pbRidings[i] = pbRiding
	}
	return &pb.ListRidingsResponse{
		Ridings: pbRidings,
	}, nil
}
