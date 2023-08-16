package user

import (
	"strings"
	"time"

	"github.com/dimiantoni/api-clean-archi-go/domain/entity"
)

// Service  interface
type Service struct {
	repo Repository
}

// NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateUser Create an user
func (s *Service) CreateUser(name, email, password, address string, age int8) (entity.ID, error) {
	e, err := entity.NewUser(name, email, password, address, age)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

// GetUser Get an user
func (s *Service) GetUser(id entity.ID) (*entity.User, error) {
	return s.repo.Get(id)
}

// SearchUsers Search users
func (s *Service) SearchUsers(query string) ([]*entity.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListUsers List users
func (s *Service) ListUsers() ([]*entity.User, error) {
	return s.repo.List()
}

// DeleteUser Delete an user
func (s *Service) DeleteUser(id entity.ID) error {
	u, err := s.GetUser(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// UpdateUser Update an user
func (s *Service) UpdateUser(e *entity.User) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

func (s *Service) Login(email, password string) (*entity.User, error) {
	users, err := s.repo.SearchUserByEmail(email)

	if err != nil {
		return &entity.User{}, err
	}

	return &entity.User{
		ID:       users.ID,
		Name:     users.Name,
		Password: users.Password,
		Email:    users.Email,
		Address:  users.Address,
		Age:      users.Age,
	}, nil
}
