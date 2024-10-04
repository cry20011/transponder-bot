package usecase

import "context"

type Users struct {
	users UsersRepo
}

func NewUsers(u UsersRepo) *Users {
	return &Users{users: u}
}

func (u *Users) AddUser(ctx context.Context, username string) error {
	return u.users.AddUser(ctx, username)
}
