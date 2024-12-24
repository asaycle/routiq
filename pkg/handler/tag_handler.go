package handler

import (
	"context"
	"log"

	pb "github.com/asaycle/motify.git/api/proto/v1"
	"github.com/asaycle/motify.git/pkg/domain/repository"
	"github.com/asaycle/motify.git/pkg/domain/usecase"
	"github.com/asaycle/motify.git/pkg/infrastructure/db"
	"github.com/asaycle/motify.git/pkg/server/handlers"
	"google.golang.org/grpc"
)

func init() {
	handlers.RegisterHandler(NewTagHandler())
}

type TagHandler struct {
	pb.UnimplementedTagServiceServer

	useCase usecase.TagUsecase
}

func NewTagHandler() *TagHandler {
	pgdb, err := db.NewPgDB("localhost", 5432, "root", "root", "motify")
	if err != nil {
		log.Panic("failed initialize pgdb", err)
	}
	txManager := repository.NewTransactionManager(pgdb)
	repo := repository.NewTagRepositoryImpl()
	return &TagHandler{
		useCase: usecase.NewTagUsecaseImpl(repo, txManager),
	}
}

func (h *TagHandler) Register(s *grpc.Server) {
	pb.RegisterTagServiceServer(s, h)
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
