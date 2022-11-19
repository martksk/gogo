package service

import "github.com/GDSC-KMUTT/totp-session/repository"

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) userService {
	return userService{repo: repo}
}

func (s userService) SignUp(email string, password string) (*int64, *string, *string, error) {
	s.repo.CreateUser(email, password, "")
	return nil, nil, nil, nil
}

func (s userService) SignIn(email string, password string) (*int64, error) {
	s.repo.CheckUser(email)
	return nil, nil
}