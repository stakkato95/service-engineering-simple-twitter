package service

type UserService interface {
}

type defaultUserService struct {
}

func NewUserService() UserService {
	return &defaultUserService{}
}
