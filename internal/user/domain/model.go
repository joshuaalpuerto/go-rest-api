package userdomain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrNameNotValid = errors.New("name is not valid")
)

// represent the company entity in the domain
type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}

type UserDB struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedBy string    `db:"created_by"`
	UpdatedBy string    `db:"updated_by"`
}

func (u UserDB) ToDomain() (User, error) {
	// we can use type conversation because properties are the same
	return User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		CreatedBy: u.CreatedBy,
		UpdatedBy: u.UpdatedBy,
	}, nil
}

func ToUserEntities(dbs []UserDB) ([]User, error) {
	users := make([]User, len(dbs))
	for i, db := range dbs {
		user, err := db.ToDomain()
		if err != nil {
			return nil, err
		}
		users[i] = user
	}
	return users, nil
}

type NewUser struct {
	Name      string
	Email     string
	Password  string
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ToDomainCompany converts PostCompany to domain Company
func (p *NewUser) ToDomainEntity(userId string) (NewUser, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return NewUser{}, fmt.Errorf("failed to hash password: %w", err)
	}

	return NewUser{
		Name:      p.Name,
		Email:     p.Email,
		Password:  string(hashedPassword),
		CreatedBy: uuid.MustParse(userId),
		UpdatedBy: uuid.MustParse(userId),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
