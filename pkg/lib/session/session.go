package session

import "context"

type Role int

const (
	RoleGuest Role = iota
	RoleUser
	RoleAdmin
)

var roleMap map[string]Role = map[string]Role{
	"":      RoleGuest,
	"Guest": RoleGuest,
	"User":  RoleUser,
	"Admin": RoleAdmin,
}

type Session struct {
	UserID string
	Role   Role
}

type (
	ctxKeyUID  struct{}
	ctxKeyRole struct{}
)

func GetSession(ctx context.Context) *Session {
	uid := ctx.Value(ctxKeyUID{}).(string)
	role := ctx.Value(ctxKeyRole{}).(Role)
	return &Session{
		UserID: uid,
		Role:   role,
	}
}

func SetGuestRole(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, ctxKeyUID{}, "GUEST")
	ctx = context.WithValue(ctx, ctxKeyRole{}, RoleGuest)
	return ctx
}

func SetUserRole(ctx context.Context, uid string) context.Context {
	ctx = context.WithValue(ctx, ctxKeyUID{}, uid)
	ctx = context.WithValue(ctx, ctxKeyRole{}, RoleUser)
	return ctx
}
