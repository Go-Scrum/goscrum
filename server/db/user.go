package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	// given_name
	FirstName string `json:"first_name"`
	// family_name
	LastName string `json:"last_name"`
	// email
	Email string `json:"email"`
}

// BeforeUpdate is a hook to set the created_at column to UNIX timestamp int.
func (m *User) BeforeUpdate(scope *gorm.Scope) error {
	return scope.SetColumn("UpdatedAt", time.Now())
}

// BeforeCreate is a hook to set the created_at column to UNIX timestamp int.
func (m *User) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("UpdatedAt", time.Now())

	if err != nil {
		return err
	}

	return scope.SetColumn("CreatedAt", time.Now())
}

func NewUser(firstName, lastName, email string) User {
	return User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}
