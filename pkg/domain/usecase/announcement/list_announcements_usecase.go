package announcement

import (
	"context"

	"github.com/asaycle/routiq/pkg/domain/entity"
	"github.com/asaycle/routiq/pkg/domain/repository"
	"golang.org/x/xerrors"
)

type ListAnnouncementsUsecase interface {
	Exec(ctx context.Context) ([]*entity.Announcement, error)
}

func NewListAnnouncementsUsecase(repo repository.AnnouncementRepository) ListAnnouncementsUsecase {
	return &listAnnouncementsUsecase{repo: repo}
}

type listAnnouncementsUsecase struct {
	repo repository.AnnouncementRepository
}

func (u *listAnnouncementsUsecase) Exec(ctx context.Context) ([]*entity.Announcement, error) {
	anns, err := u.repo.ListAnnouncements(ctx, nil)
	if err != nil {
		return nil, xerrors.Errorf("Error ListAnnouncements: %w", err)
	}
	return anns, nil
}
