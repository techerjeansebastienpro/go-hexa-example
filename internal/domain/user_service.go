package domain

import (
	"errors"
)

type UserService struct {
	output UsersOuput
}

func NewUserService(out UsersOuput) *UserService {
	return &UserService{output: out}
}

func (s *UserService) FindOne(request FindOneRequest) (User, error) {
	if request.ID != "" {
		return s.output.FindOneById(FindOneById{ID: request.ID})
	}
	if request.Email != "" {
		return s.output.FindOneByEmail(FindOneByEmail{Email: request.Email})
	}
	return User{}, errors.New("invalid request")
}

func (s *UserService) Insert(request CreateUser) (User, error) {
	return s.output.InsertOne(request)
}
