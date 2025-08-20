package entity

import (
	pb "github.com/asaycle/routiq/api/proto/v1"
)

type City struct {
	ID          string
	PrefID      string
	DisplayName string
}

func NewCity(id string, prefID string, displayName string) *City {
	return &City{
		ID:          id,
		PrefID:      prefID,
		DisplayName: displayName,
	}
}

func (c *City) ToProto() *pb.City {
	return &pb.City{
		Name:        pb.FormatCityName(c.PrefID, c.ID),
		DisplayName: c.DisplayName,
		Pref:        pb.FormatPrefName(c.PrefID),
	}
}
