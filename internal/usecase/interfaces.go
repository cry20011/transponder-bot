package usecase

import "context"

type (
	UsersRepo interface {
		AddUser(ctx context.Context, username string) error
	}
)
