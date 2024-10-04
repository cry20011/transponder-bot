package repo

import (
	"context"
	"transponder-bot/pkg/postgres"
)

type Users struct {
	postgres *postgres.Postgres
}

func NewUsers(p *postgres.Postgres) *Users {
	return &Users{
		postgres: p,
	}
}

func (u *Users) AddUser(ctx context.Context, username string) error {
	insertUserStmt := `insert into "users"("name") values($1)`

	_, err := u.postgres.ExecContext(ctx, insertUserStmt, username)
	return err
}
