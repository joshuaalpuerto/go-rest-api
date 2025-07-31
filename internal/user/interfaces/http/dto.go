package userhttp

import (
	"time"

	userdomain "github.com/joshuaalpuerto/go-rest-api/internal/user/domain"
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
