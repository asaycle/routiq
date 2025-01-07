package repository

import (
	"time"

	"github.com/asaycle/routiq.git/pkg/domain/entity"
)

type Riding struct {
	ID        string    `db:"id"`
	RouteID   string    `db:"route_id"`
	Title     string    `db:"title"`
	Note      string    `db:"note"`
	Score     int       `db:"score"`
	Date      time.Time `db:"date"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func fromRidingEntity(riding *entity.Riding) (*Riding, error) {
	return &Riding{
		ID:      riding.ID,
		RouteID: riding.RouteID,
		Title:   riding.Title,
		Date:    riding.Date,
		Score:   riding.Score,
		Note:    riding.Note,
	}, nil
}

func toRidingEntity(model *Riding) (*entity.Riding, error) {
	return &entity.Riding{
		ID:    model.ID,
		Title: model.Title,
		Note:  model.Note,
	}, nil
}
