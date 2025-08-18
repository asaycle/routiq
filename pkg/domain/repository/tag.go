package repository

import "github.com/asaycle/routiq/pkg/domain/entity"

type Tag struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func fromTagEntity(e *entity.Tag) *Tag {
	return &Tag{
		ID:   e.ID,
		Name: e.DisplayName,
	}
}

func toTagEntity(model *Tag) *entity.Tag {
	return &entity.Tag{
		ID:          model.ID,
		DisplayName: model.Name,
	}
}
