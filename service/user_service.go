package service

import "mock.com/repository"

type UserService struct {
	ur repository.IUserRepository
}

type IUserService interface {
	Create() error
}

func NewUserService(ur repository.IUserRepository) IUserService {
	return &UserService{
		ur: ur,
	}
}

func (s *UserService) Create() error {
	return s.ur.CreateUser()
}
