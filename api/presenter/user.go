package presenter

import (
	"github.com/dimiantoni/api-clean-archi-go/domain/entity"
)

// User data
type User struct {
	ID       entity.ID `json:"id" bson:"id",omitempty`
	Name     string    `json:"name" bson:"name",omitempty`
	Email    string    `json:"email" bson:"email",omitempty`
	Password string    `json:"password" bson:"password",omitempty`
	Address  string    `json:"address" bson:"address",omitempty`
	Age      int8      `json:"age" bson:"age",omitempty`
}

type AuthToken struct {
	Token string `json:"token" bson:"token",omitempty`
}
