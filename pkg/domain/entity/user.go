package entity

import "github.com/rs/xid"

type User struct {
	ID  string
	Sub string
}

func NewUser(subject string) *User {
	return &User{
		ID:  xid.New().String(),
		Sub: subject,
	}
}

type UserProfile struct {
	ID        string
	UserID    string
	Bio       string
	AvatarURL string
}

func NewUserProfile(userID string, bio string, avatarURL string) *UserProfile {
	return &UserProfile{
		ID:        xid.New().String(),
		UserID:    userID,
		Bio:       bio,
		AvatarURL: avatarURL,
	}
}
