package userusecases

import (
	"context"
	"fmt"

	userdomain "github.com/joshuaalpuerto/go-rest-api/internal/user/domain"
)

func (c *UserUsecase) CreateUser(nc userdomain.NewUser) (*userdomain.User, error) {
	// this should come from request
	userId := "c8d9c08f-4f87-4c5c-8862-2f4abac75f1f"

	// Validate if incoming payload adhere to the domain model
	userEntity, err := nc.ToDomainEntity(userId)
	if err != nil {
		return nil, fmt.Errorf("Invalid user: [%+v]: %w", nc, err)
	}

	result, err := c.userRepository.Create(context.Background(), userEntity)
	if err != nil {
		return nil, err
	}

	user, err := result.ToDomain()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
