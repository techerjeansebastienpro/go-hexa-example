package domain

type UsersInput interface {
	GetByID(ID string) (User, error)
	CreateOne(request CreateUser) (User, error)
}

type UsersOuput interface {
	FindOneById(request FindOneById) (User, error)
	FindOneByEmail(request FindOneByEmail) (User, error)
	InsertOne(request CreateUser) (User, error)
}
