package handler

import (
	"context"
	"log"
	"time"

	pb "github.com/asaycle/routiq/api/proto/v1"
	"github.com/asaycle/routiq/pkg/domain/repository"
	"github.com/asaycle/routiq/pkg/domain/usecase"
	"github.com/asaycle/routiq/pkg/infrastructure/db"
	"github.com/asaycle/routiq/pkg/lib/config"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ActivateTouringHandler(s *grpc.Server, cfg *config.Config) {
	pb.RegisterTouringServiceServer(s, NewTouringHandler(cfg))
}

type TouringHandler struct {
	pb.UnimplementedTouringServiceServer
	cfg     *config.Config
	useCase usecase.TouringUsecase
}

func NewTouringHandler(cfg *config.Config) *TouringHandler {
	pgdb, err := db.NewPgDBFromCfg(&cfg.DB)
	if err != nil {
		log.Panic("failed initialize pgdb", err)
	}
	txManager := repository.NewTransactionManager(pgdb)
	return &TouringHandler{
		cfg: cfg,
		useCase: usecase.NewTouringUsecaseImpl(
			repository.NewTouringRepositoryImpl(),
			repository.NewRouteRepositoryImpl(),
			txManager,
		),
	}
}

func (h *TouringHandler) CreateTouring(ctx context.Context, req *pb.CreateTouringRequest) (*pb.Touring, error) {
	routeName, err := pb.ParseRouteName(req.Parent)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid parent route name")
	}
	date := time.Date(int(req.Touring.Date.Year), time.Month(req.Touring.Date.Month), int(req.Touring.Date.Day), 0, 0, 0, 0, time.Local)
	tags := make([]string, len(req.Touring.Tags))
	for i, tagNameStr := range req.Touring.Tags {
		tagName, err := pb.ParseTagName(tagNameStr)
		if err != nil {
			return nil, xerrors.Errorf("Error ParseTagName: %w", err)
		}
		tags[i] = tagName.GetTag()
	}
	touring, err := h.useCase.CreateTouring(
		ctx,
		routeName.GetRoute(),
		req.Touring.DisplayName,
		date,
		req.Touring.Note,
		int(req.Touring.Score),
		tags,
	)
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
