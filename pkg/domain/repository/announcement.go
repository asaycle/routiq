package repository

import (
	"time"

	"github.com/asaycle/routiq/pkg/domain/entity"
	"golang.org/x/xerrors"
)

type Announcement struct {
	ID        int    `yaml:"id"`
	Title     string `yaml:"title"`
	Body      string `yaml:"body"`
	CreatedAt string `yaml:"created_at"`
}

func (a *Announcement) toEntity() (*entity.Announcement, error) {
	t, err := time.Parse("2006-01-02", a.CreatedAt)
	if err != nil {
		// フォーマットエラーの場合はスキップ or エラー返却にする
		return nil, xerrors.Errorf("invalid date %q: %w", a.CreatedAt, err)
	}
	return entity.NewAnnouncement(a.ID, a.Title, a.Body, t), nil
}
