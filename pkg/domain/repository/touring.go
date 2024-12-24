package repository

import (
	"time"

	"github.com/asaycle/motify.git/pkg/domain/entity"
)

type Touring struct {
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

func fromTouringEntity(touring *entity.Touring) (*Touring, error) {
	return &Touring{
		ID:      touring.ID,
		RouteID: touring.RouteID,
		Title:   touring.Title,
		Date:    touring.Date,
		Score:   touring.Score,
		Note:    touring.Note,
	}, nil
}

func toTouringEntity(model *Touring) (*entity.Touring, error) {
	return &entity.Touring{
		ID:    model.ID,
		Title: model.Title,
		Note:  model.Note,
	}, nil
}
