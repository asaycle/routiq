package repository

import "github.com/asaycle/routiq.git/pkg/domain/entity"

type Tag struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func fromTagEntity(e *entity.Tag) *Tag {
	return &Tag{
		ID:   e.ID,
		Name: e.Name,
	}
}

func toTagEntity(model *Tag) *entity.Tag {
	return &entity.Tag{
		ID:   model.ID,
		Name: model.Name,
	}
}
