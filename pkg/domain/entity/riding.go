package entity

import (
	"time"

	pb "github.com/asaycle/routiq.git/api/proto/v1"
	"github.com/rs/xid"
)

type Riding struct {
	ID      string
	RouteID string
	Title   string
	Date    time.Time
	Note    string
	Score   int
	Tags    []*Tag
	UserID  string
}

func (t *Riding) ToProto() (*pb.Riding, error) {
	return &pb.Riding{
		Id:    t.ID,
		Title: t.Title,
	}, nil
}

func NewRiding(routeID string, title string, date time.Time, note string, score int, tags []*Tag) *Riding {
	return &Riding{
		ID:      xid.New().String(),
		RouteID: routeID,
		Title:   title,
		Date:    date,
		Note:    note,
		Score:   score,
		Tags:    tags,
		UserID:  "admin",
	}
}

type RidingTag struct {
	ID        string
	RidingID  string
	TagID     string
	CreatedAt time.Time
}

func NewRidingTag(ridingID string, tagID string) *RidingTag {
	return &RidingTag{
		ID:       xid.New().String(),
		RidingID: ridingID,
		TagID:    tagID,
	}
}
