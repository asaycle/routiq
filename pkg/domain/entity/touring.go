package entity

import (
	"time"

	pb "github.com/asaycle/routiq/api/proto/v1"
	"github.com/rs/xid"
)

type Touring struct {
	ID          string
	RouteID     string
	DisplayName string
	Date        time.Time
	Note        string
	Score       int
	Tags        []*Tag
	UserID      string
}

func (t *Touring) ToProto() (*pb.Touring, error) {
	return &pb.Touring{
		Name:        pb.FormatTouringName(t.RouteID, t.ID),
		DisplayName: t.DisplayName,
	}, nil
}

func NewTouring(routeID string, displayName string, date time.Time, note string, score int, tags []*Tag) *Touring {
	return &Touring{
		ID:          xid.New().String(),
		RouteID:     routeID,
		DisplayName: displayName,
		Date:        date,
		Note:        note,
		Score:       score,
		Tags:        tags,
		UserID:      "admin",
	}
}

type TouringTag struct {
	ID        string
	TouringID string
	TagID     string
	CreatedAt time.Time
}

func NewTouringTag(touringID string, tagID string) *TouringTag {
	return &TouringTag{
		ID:        xid.New().String(),
		TouringID: touringID,
		TagID:     tagID,
	}
}
