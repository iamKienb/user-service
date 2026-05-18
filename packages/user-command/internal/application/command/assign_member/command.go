package assign_member

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type User struct {
	ID   shared.UserID
	Name string
}

type MemberRole struct {
	ID      shared.UserID
	Name    string
	RoleIDs []shared.RoleID
}

type Command struct {
	User        User
	ShopID      shared.ShopID
	MemberRoles []MemberRole
	Action      string
}

type Result struct {
	Success bool
}

type Executor interface {
	Execute(ctx context.Context, cmd Command) (*Result, error)
}
