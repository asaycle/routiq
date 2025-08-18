package repository

import "github.com/asaycle/routiq/pkg/domain/entity"

type User struct {
	ID  string `db:"id"`
	Sub string `db:"google_sub"`
}

func fromUserEntity(user *entity.User) *User {
	return &User{
		ID:  user.ID,
		Sub: user.Sub,
	}
}

func toUserEntity(model *User) *entity.User {
	return &entity.User{
		ID:  model.ID,
		Sub: model.Sub,
	}
}
