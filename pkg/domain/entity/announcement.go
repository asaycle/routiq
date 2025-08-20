package entity

import (
	"strconv"
	"time"

	pb "github.com/asaycle/routiq/api/proto/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Announcement struct {
	ID         int
	Title      string
	Body       string
	CreateTime time.Time
}

func NewAnnouncement(id int, title, body string, createTime time.Time) *Announcement {
	return &Announcement{
		ID:         id,
		Title:      title,
		Body:       body,
		CreateTime: createTime,
	}
}

func (a *Announcement) ToProto() *pb.Announcement {
	return &pb.Announcement{
		Name:       pb.FormatAnnouncementName(strconv.Itoa(a.ID)),
		Title:      a.Title,
		Body:       a.Body,
		CreateTime: timestamppb.New(a.CreateTime),
	}
}
