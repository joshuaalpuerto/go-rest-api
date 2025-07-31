package userusecases

import (
	"context"

	userdomain "github.com/joshuaalpuerto/go-rest-api/internal/api/user/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]userdomain.UserDB, error)
	FindOneByID(ctx context.Context, id string) (*userdomain.UserDB, error)
	Create(ctx context.Context, user userdomain.NewUser) (*userdomain.UserDB, error)
	Update(ctx context.Context, user userdomain.User) (*userdomain.UserDB, error)
	Delete(ctx context.Context, id string) (*userdomain.UserDB, error)
}

type UserUsecase struct {
	userRepository UserRepository
}

func NewUserUsecase(userRepository UserRepository) UserUsecase {
	return UserUsecase{
		userRepository: userRepository,
	}
}
