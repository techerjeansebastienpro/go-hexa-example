package domain

type User struct {
	ID    string
	Email string
}

type FindOneById struct {
	ID string
}

type FindOneByEmail struct {
	Email string
}

type FindOneRequest struct {
	FindOneById
	FindOneByEmail
}

type CreateUser struct {
	Email    string
	Password string
}
