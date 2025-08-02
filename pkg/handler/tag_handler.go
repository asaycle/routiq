package handler

import (
	"context"
	"log"
	"os"

	pb "github.com/asaycle/routiq.git/api/proto/v1"
	"github.com/asaycle/routiq.git/pkg/domain/repository"
	"github.com/asaycle/routiq.git/pkg/domain/usecase"
	"github.com/asaycle/routiq.git/pkg/infrastructure/db"
	"github.com/asaycle/routiq.git/pkg/lib/config"
	"google.golang.org/grpc"
)

func ActivateTagHandler(s *grpc.Server, cfg *config.Config) {
	pb.RegisterTagServiceServer(s, NewTagHandler(cfg))
}

type TagHandler struct {
	pb.UnimplementedTagServiceServer
	cfg     *config.Config
	useCase usecase.TagUsecase
}

func NewTagHandler(cfg *config.Config) *TagHandler {
	pgdb, err := db.NewPgDB(
		os.Getenv("POSTGRES_HOST"),
		5432,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	if err != nil {
		log.Panic("failed initialize pgdb", err)
	}
	txManager := repository.NewTransactionManager(pgdb)
	repo := repository.NewTagRepositoryImpl()
	return &TagHandler{
		cfg:     cfg,
		useCase: usecase.NewTagUsecaseImpl(repo, txManager),
	}
}

func (h *TagHandler) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.Tag, error) {
	tag, err := h.useCase.CreateTag(ctx, req.Tag.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Tag{
		Id:   tag.ID,
		Name: tag.Name,
	}, nil
}

func (h *TagHandler) ListTags(ctx context.Context, req *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
	log.Printf("hello list tags")
	tags, err := h.useCase.ListTags(ctx)
	if err != nil {
		return nil, err
	}
	pbTags := make([]*pb.Tag, len(tags))
	for i, tag := range tags {
		pbTags[i] = &pb.Tag{
			Id:   tag.ID,
			Name: tag.Name,
		}
	}
	return &pb.ListTagsResponse{Tags: pbTags}, nil
}

var _ pb.TagServiceServer = &TagHandler{}
