package postgres

import "fmt"

type Options struct {
	Host, Port     string
	User, Password string
	DbName         string
	SslMode        string
}

func (o Options) makeConnStr() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		o.Host, o.Port, o.User, o.Password, o.DbName, o.SslMode)
}
