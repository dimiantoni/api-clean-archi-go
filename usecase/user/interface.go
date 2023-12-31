package user

import (
	"github.com/dimiantoni/api-clean-archi-go/domain/entity"
)

// Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.User, error)
	SearchUserByEmail(email string) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	List() ([]*entity.User, error)
}

// Writer user writer
type Writer interface {
	Create(e *entity.User) (entity.ID, error)
	Update(e *entity.User) error
	Delete(id entity.ID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	GetUser(id entity.ID) (*entity.User, error)
	SearchUsers(query string) ([]*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(name, email, password, address string, age int8) (entity.ID, error)
	UpdateUser(*entity.User) error
	DeleteUser(id entity.ID) error
	Login(email, password string) (*entity.User, error)
}
