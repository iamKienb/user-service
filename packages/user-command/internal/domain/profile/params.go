package profile

import "user-command-module/internal/domain/shared"

type NewProfileParams struct {
	UserID   shared.UserID
	FullName string
	Gender   GenderEnum
}
