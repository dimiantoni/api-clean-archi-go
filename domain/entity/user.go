package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User data
type User struct {
	ID        ID
	Name      string
	Email     string
	Password  string
	Address   string
	Age       int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser create a new user
func NewUser(name, email, password, address string, age int8) (*User, error) {
	u := &User{
		ID:        NewID(),
		Name:      name,
		Email:     email,
		Password:  password,
		Address:   address,
		Age:       age,
		CreatedAt: time.Now(),
	}
	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	err = u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

// Validate validate data
func (u *User) Validate() error {
	if u.Email == "" || u.Name == "" || u.Address == "" {
		return ErrInvalidEntity
	}

	return nil
}

// ValidatePassword validate user password
func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
