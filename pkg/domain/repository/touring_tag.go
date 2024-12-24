package repository

import (
	"time"

	"github.com/asaycle/motify.git/pkg/domain/entity"
)

type TouringTag struct {
	ID        string    `db:"id"`
	TouringID string    `db:"touring_id"`
	TagID     string    `db:"tag_id"`
	CreatedAt time.Time `db:"created_at"`
}

func fromTouringTagEntity(e *entity.TouringTag) (*TouringTag, error) {
	return &TouringTag{
		ID:        e.ID,
		TouringID: e.TouringID,
		TagID:     e.TagID,
		CreatedAt: e.CreatedAt,
	}, nil
}
