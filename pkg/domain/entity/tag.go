package entity

import "github.com/rs/xid"

type Tag struct {
	ID          string
	DisplayName string
}

func NewTag(name string) *Tag {
	return &Tag{
		ID:          xid.New().String(),
		DisplayName: name,
	}
}
