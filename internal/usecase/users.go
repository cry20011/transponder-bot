package usecase

type Users struct {
	users UsersRepo
}

func NewUsers(u UsersRepo) *Users {
	return &Users{users: u}
}
