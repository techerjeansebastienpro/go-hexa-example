package infrastructure

import (
	"context"

	"github.com/techerjeansebastienpro/go-hexa-example/internal/domain"
	db "github.com/techerjeansebastienpro/go-hexa-example/pkg/models"
)

type UserRepository struct {
	prisma *db.PrismaClient
}

func NewUserRepository(prisma *db.PrismaClient) *UserRepository {
	return &UserRepository{
		prisma: prisma,
	}
}

func (r *UserRepository) FindOneById(request domain.FindOneById) (domain.User, error) {
	ctx := context.Background()
	foundUser, err := r.prisma.User.FindUnique(
		db.User.ID.Equals(request.ID),
	).Exec(ctx)

	return domain.User{
		ID: foundUser.ID,
	}, err
}

func (r *UserRepository) FindOneByEmail(request domain.FindOneByEmail) (domain.User, error) {
	ctx := context.Background()
	foundUser, err := r.prisma.User.FindUnique(
		db.User.Email.Equals(request.Email),
	).Exec(ctx)

	return domain.User{
		ID: foundUser.ID,
	}, err
}

func (r *UserRepository) InsertOne(request domain.CreateUser) (domain.User, error) {
	ctx := context.Background()
	createdUser, err := r.prisma.User.CreateOne(
		db.User.Email.Set(request.Email),
		db.User.Password.Set(request.Password),
	).Exec(ctx)

	return domain.User{
		ID:    createdUser.ID,
		Email: createdUser.Email,
	}, err
}
