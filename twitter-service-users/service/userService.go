package service

import (
	"net/http"

	"github.com/stakkato95/service-engineering-go-lib/errs"
	"github.com/stakkato95/twitter-service-users/domain"
	"github.com/stakkato95/twitter-service-users/jwt"
)

var passwordErr = errs.NewAppError(
	"can not authorize user: wrong password",
	http.StatusUnauthorized)

type UserService interface {
	Create(*domain.User) (string, *errs.AppError)
	Authenticate(*domain.User) (string, *errs.AppError)
}

type defaultUserService struct {
	repo domain.UserRepo
}

func NewUserService(repo domain.UserRepo) UserService {
	return &defaultUserService{repo}
}

func (s *defaultUserService) Create(user *domain.User) (string, *errs.AppError) {
	if err := s.repo.Create(user); err != nil {
		return "", errs.NewAppError(
			"can not create user: "+err.Error(),
			http.StatusInternalServerError)
	}

	return generateToken(user.Username)
}

func (s *defaultUserService) Authenticate(user *domain.User) (string, *errs.AppError) {
	hashedPassword, err := s.repo.Authenticate(user)
	if err != nil {
		return "", errs.NewAppError(
			"can not authorize user: "+err.Error(),
			http.StatusUnauthorized)
	}

	if ok := checkPasswordHash(user.Password, hashedPassword); !ok {
		return "", passwordErr
	}

	return generateToken(user.Username)
}

func checkPasswordHash(hashedPassword, hash string) bool {
	return true
	// err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(hashedPassword))
	// return err == nil
}

func generateToken(username string) (string, *errs.AppError) {
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", errs.NewAppError(
			"can not generate jwt token: "+err.Error(),
			http.StatusInternalServerError)
	}
	return token, nil
}
