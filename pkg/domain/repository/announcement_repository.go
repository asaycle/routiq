package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/asaycle/routiq/pkg/domain/entity"
	"golang.org/x/xerrors"
	"gopkg.in/yaml.v3"
)

type AnnouncementRepository interface {
	ListAnnouncements(ctx context.Context, tx Tx) ([]*entity.Announcement, error)
	GetAnnouncement(ctx context.Context, tx Tx, announcementID string) (*entity.Announcement, error)
}

func NewAnnouncementRepository() AnnouncementRepository {
	anns, err := loadAnnouncements("announcements.yml")
	if err != nil {
		panic("Failed to load announcements")
	}
	m := map[int]*entity.Announcement{}
	for _, ann := range anns {
		m[ann.ID] = ann
	}
	return &announcementRepository{m: m}
}

type announcementRepository struct {
	m map[int]*entity.Announcement
}

func (r *announcementRepository) ListAnnouncements(ctx context.Context, tx Tx) ([]*entity.Announcement, error) {
	var announcements []*entity.Announcement
	for _, ann := range r.m {
		announcements = append(announcements, ann)
	}
	return announcements, nil
}

func (r *announcementRepository) GetAnnouncement(ctx context.Context, tx Tx, announcementID string) (*entity.Announcement, error) {
	// ここでデータベースからアナウンスメントを取得するロジックを実装
	return nil, nil // 仮の実装
}

func loadAnnouncements(path string) ([]*entity.Announcement, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var raw []Announcement
	if err := yaml.Unmarshal(b, &raw); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	announcements := make([]*entity.Announcement, 0, len(raw))
	for i, r := range raw {

		ann, err := r.toEntity()
		if err != nil {
			return nil, xerrors.Errorf("convert to entity: %w", err)
		}
		announcements[i] = ann
	}
	return announcements, nil
}
