package entity

import (
	pb "github.com/asaycle/routiq/api/proto/v1"
)

type Pref struct {
	ID          string
	DisplayName string
}

func NewPref(id string, displayName string) *Pref {
	return &Pref{
		ID:          id,
		DisplayName: displayName,
	}
}

func (p *Pref) ToProto() *pb.Pref {
	return &pb.Pref{
		Name:        pb.FormatPrefName(p.ID),
		DisplayName: p.DisplayName,
	}
}
