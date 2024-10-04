package usecase

type (
	UsersRepo interface {
		AddUser(username string) error
	}
)
