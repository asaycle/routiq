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
	"github.com/asaycle/routiq.git/pkg/server/handlers"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

func init() {
	handlers.RegisterHandler(NewTouringHandler())
}

type TouringHandler struct {
	pb.UnimplementedTouringServiceServer

	useCase usecase.TouringUsecase
}

func NewTouringHandler() *TouringHandler {
	pgdb, err := db.NewPgDB("localhost", 5432, "root", "root", "routiq")
	if err != nil {
		log.Panic("failed initialize pgdb", err)
	}
	txManager := repository.NewTransactionManager(pgdb)
	return &TouringHandler{
		useCase: usecase.NewTouringUsecaseImpl(
			repository.NewTouringRepositoryImpl(),
			repository.NewRouteRepositoryImpl(),
			txManager,
		),
	}
}

func (h *TouringHandler) Register(s *grpc.Server) {
	pb.RegisterTouringServiceServer(s, h)
}

func (h *TouringHandler) CreateTouring(ctx context.Context, req *pb.CreateTouringRequest) (*pb.Touring, error) {
	date := time.Date(int(req.Touring.Date.Year), time.Month(req.Touring.Date.Month), int(req.Touring.Date.Day), 0, 0, 0, 0, time.Local)
	tags := make([]*entity.Tag, len(req.Touring.Tags))
	for i, tag := range req.Touring.Tags {
		tags[i] = &entity.Tag{
			ID:   tag.Id,
			Name: tag.Name,
		}
	}
	touring, err := h.useCase.CreateTouring(ctx, req.Touring.RouteId, req.Touring.Title, date, req.Touring.Note, int(req.Touring.Score), tags)
	if err != nil {
		return nil, xerrors.Errorf("Error CreateTouring: %w", err)
	}
	pbTouring, err := touring.ToProto()
	if err != nil {
		return nil, xerrors.Errorf("Error ToProto: %w", err)
	}
	return pbTouring, nil
}

func (h *TouringHandler) ListTourings(ctx context.Context, req *pb.ListTouringsRequest) (*pb.ListTouringsResponse, error) {
	tourings, err := h.useCase.ListTourings(ctx, req.Filter)
	if err != nil {
		return nil, xerrors.Errorf("Error ListTourings: %w", err)
	}
	pbTourings := make([]*pb.Touring, len(tourings))
	for i, e := range tourings {
		pbTouring, err := e.ToProto()
		if err != nil {
			return nil, xerrors.Errorf("Error ToProto: %w", err)
		}
		pbTourings[i] = pbTouring
	}
	return &pb.ListTouringsResponse{
		Tourings: pbTourings,
	}, nil
}
