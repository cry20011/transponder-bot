package repo

import "transponder-bot/pkg/postgres"

type Users struct {
	postgres *postgres.Postgres
}

func New(p *postgres.Postgres) Users {
	return Users{
		postgres: p,
	}
}

func (u *Users) AddUser(username string) error {
	insertUserStmt := `insert into "users"("name") values($1)`

	_, err := u.postgres.Exec(insertUserStmt, username)
	return err
}
