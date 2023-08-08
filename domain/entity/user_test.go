package entity_test

import (
	"testing"

	"github.com/dimiantoni/api-clean-archi-go/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	u, err := entity.NewUser("Bill", "bill@google.com", "123456", "Lime street 129", 18)
	assert.Nil(t, err)
	assert.Equal(t, u.Name, "Bill")
	assert.NotNil(t, u.ID)
	assert.NotEqual(t, u.Password, "new_password")
}

func TestValidatePassword(t *testing.T) {
	u, _ := entity.NewUser("Bill", "bill@google.com", "new_password", "Lime street 129", 18)
	err := u.ValidatePassword("new_password")
	assert.Nil(t, err)
	err = u.ValidatePassword("wrong_password")
	assert.NotNil(t, err)

}

func TestUserValidate(t *testing.T) {
	type test struct {
		name     string
		email    string
		password string
		address  string
		age      int8
		want     error
	}

	tests := []test{
		{
			name:     "Bill",
			email:    "bill@google.com",
			password: "new_password",
			address:  "Lime street 129",
			age:      18,
			want:     nil,
		},
		// {
		// 	name:     "Bill",
		// 	email:    "",
		// 	password: "new_password",
		// 	address:  "Lime street 129",
		// 	age:      18,
		// 	want:     entity.ErrInvalidEntity,
		// },
		// {
		// 	name:     "Bill",
		// 	email:    "bill@google.com",
		// 	password: "",
		// 	address:  "Lime street 129",
		// 	age:      18,
		// 	want:     entity.ErrInvalidEntity,
		// },
		// {
		// 	name:     "Bill",
		// 	email:    "bill@google.com",
		// 	password: "new_password",
		// 	address:  "",
		// 	age:      18,
		// 	want:     entity.ErrInvalidEntity,
		// },
	}
	for _, tc := range tests {

		_, err := entity.NewUser(tc.name, tc.email, tc.password, tc.address, tc.age)
		assert.Equal(t, err, tc.want)
	}

}
