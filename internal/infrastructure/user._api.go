package infrastructure

import (
	"github.com/techerjeansebastienpro/go-hexa-example/internal/domain"
)

type UserApi struct {
	userService domain.UserService
}

func NewUserApi(userService domain.UserService) *UserApi {
	return &UserApi{
		userService: userService,
	}
}

func (a *UserApi) GetByID(ID string) (domain.User, error) {

	return a.userService.FindOne(domain.FindOneRequest{
		FindOneById: domain.FindOneById{
			ID: ID,
		},
	})
}

func (a *UserApi) CreateOne(request domain.CreateUser) (domain.User, error) {
	return a.userService.Insert(domain.CreateUser{
		Email:    request.Email,
		Password: request.Password,
	})
}
