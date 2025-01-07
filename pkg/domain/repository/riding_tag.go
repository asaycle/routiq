package repository

import (
	"time"

	"github.com/asaycle/routiq.git/pkg/domain/entity"
)

type RidingTag struct {
	ID        string    `db:"id"`
	RidingID  string    `db:"riding_id"`
	TagID     string    `db:"tag_id"`
	CreatedAt time.Time `db:"created_at"`
}

func fromRidingTagEntity(e *entity.RidingTag) (*RidingTag, error) {
	return &RidingTag{
		ID:        e.ID,
		RidingID:  e.RidingID,
		TagID:     e.TagID,
		CreatedAt: e.CreatedAt,
	}, nil
}
