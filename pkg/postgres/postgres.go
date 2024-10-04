package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	*sql.DB
}

func New(options Options) (*Postgres, error) {
	db, err := sql.Open("postgres", options.makeConnStr())
	if err != nil {
		return nil, err
	}

	return &Postgres{DB: db}, nil
}

func (p *Postgres) Close() error {
	return p.DB.Close()
}
