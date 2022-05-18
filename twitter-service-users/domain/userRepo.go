package domain

type UserRepo interface {
	Create(*User) error
	Authenticate(*User) (string, error)
}

type defaultUserRepo struct {
}

func NewUserRepo() UserRepo {
	return &defaultUserRepo{}
}

func (r *defaultUserRepo) Create(user *User) error {
	return nil
}

func (r *defaultUserRepo) Authenticate(user *User) (string, error) {
	return "password", nil
}
