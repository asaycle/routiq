package entity

import "github.com/rs/xid"

type Tag struct {
	ID   string
	Name string
}

func NewTag(name string) *Tag {
	return &Tag{
		ID:   xid.New().String(),
		Name: name,
	}
}
