package verify_permission

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type Command struct {
	ShopID shared.ShopID
	UserID shared.UserID
	Action string
}

type Result struct {
	IsAllowed    bool
	ErrorMessage error
}

type Executor interface {
	Execute(ctx context.Context, cmd Command) (*Result, error)
}
