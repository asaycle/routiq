package handler

import (
	"context"

	pb "github.com/asaycle/routiq/api/proto/v1"
	"github.com/asaycle/routiq/pkg/domain/repository"
	usecase "github.com/asaycle/routiq/pkg/domain/usecase/announcement"
	"github.com/asaycle/routiq/pkg/lib/config"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

type AnnouncementHandler struct {
	pb.UnimplementedRouteServiceServer
	cfg     *config.Config
	useCase usecase.ListAnnouncementsUsecase
}

func ActivateAnnouncementHandler(s *grpc.Server, cfg *config.Config) {
	pb.RegisterRouteServiceServer(s, NewAnnouncementHandler(cfg))
}

func NewAnnouncementHandler(cfg *config.Config) *AnnouncementHandler {
	return &AnnouncementHandler{
		cfg:     cfg,
		useCase: usecase.NewListAnnouncementsUsecase(repository.NewAnnouncementRepository()),
	}
}

func (h *AnnouncementHandler) ListAnnouncements(ctx context.Context, req *pb.ListAnnouncementsRequest) (*pb.ListAnnouncementsResponse, error) {
	anns, err := h.useCase.Exec(ctx)
	if err != nil {
		return nil, xerrors.Errorf("Error ListAnnouncements: %v", err)
	}

	pbAnns := make([]*pb.Announcement, len(anns))
	for i, ann := range anns {
		pbAnns[i] = ann.ToProto()
	}

	return &pb.ListAnnouncementsResponse{Announcements: pbAnns}, nil
}
