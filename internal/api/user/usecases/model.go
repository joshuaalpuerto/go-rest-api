package userusecases

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	userdomain "github.com/joshuaalpuerto/go-rest-api/internal/api/user/domain"
)

// Contract of what company looks like to the client
type UserDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
}

func ToAppUser(c userdomain.User) UserDTO {
	return UserDTO{
		ID:        c.ID.String(),
		Name:      c.Name,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
		UpdatedAt: c.UpdatedAt.Format(time.RFC3339),
		CreatedBy: c.CreatedBy,
		UpdatedBy: c.UpdatedBy,
	}
}

func ToAppUsers(c []userdomain.User) []UserDTO {
	if len(c) == 0 {
		return []UserDTO{}
	}

	users := make([]UserDTO, len(c))
	for i, v := range c {
		users[i] = ToAppUser(v)
	}
	return users
}

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ToDomainCompany converts PostCompany to domain Company
func (p *NewUser) ToDomainEntity(userId string) (userdomain.NewUser, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return userdomain.NewUser{}, fmt.Errorf("failed to hash password: %w", err)
	}

	return userdomain.NewUser{
		Name:      p.Name,
		Email:     p.Email,
		Password:  string(hashedPassword),
		CreatedBy: uuid.MustParse(userId),
		UpdatedBy: uuid.MustParse(userId),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
