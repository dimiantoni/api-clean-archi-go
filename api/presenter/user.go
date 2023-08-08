package presenter

import (
	"github.com/dimiantoni/api-clean-archi-go/domain/entity"
)

// User data
type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Address  string    `json:"address"`
	Age      int8      `json:"age"`
}
