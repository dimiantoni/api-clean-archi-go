package user

import (
	"testing"
	"time"

	"github.com/dimiantoni/api-clean-archi-go/domain/entity"

	"github.com/stretchr/testify/assert"
)

func newFixtureUser() *entity.User {
	return &entity.User{
		ID:        entity.NewID(),
		Name:      "Buster",
		Email:     "buster@gmail.com",
		Password:  "123456",
		Address:   "123456",
		Age:       18,
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureUser()
	_, err := m.CreateUser(u.Name, u.Email, u.Password, u.Address, u.Age)
	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
	assert.True(t, u.UpdatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2.Name = "Lemmy"

	uID, _ := m.CreateUser(u1.Name, u1.Email, u1.Password, u1.Address, u1.Age)
	_, _ = m.CreateUser(u2.Name, u2.Email, u2.Password, u2.Address, u2.Age)

	t.Run("search", func(t *testing.T) {
		c, err := m.SearchUsers("Buster")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Buster", c[0].Name)

		c, err = m.SearchUsers("dio")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListUsers()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetUser(uID)
		assert.Nil(t, err)
		assert.Equal(t, u1.Name, saved.Name)
	})
}

func Test_Update(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureUser()
	id, err := m.CreateUser(u.Name, u.Email, u.Password, u.Address, u.Age)
	assert.Nil(t, err)
	saved, _ := m.GetUser(id)
	saved.Name = "Thor"
	assert.Nil(t, m.UpdateUser(saved))
	updated, err := m.GetUser(id)
	assert.Nil(t, err)
	assert.Equal(t, "Thor", updated.Name)
	assert.False(t, updated.UpdatedAt.IsZero())
}

func TestDelete(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newFixtureUser()
	u2 := newFixtureUser()
	u2ID, _ := m.CreateUser(u2.Name, u2.Email, u2.Password, u2.Address, u2.Age)

	err := m.DeleteUser(u1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteUser(u2ID)
	assert.Nil(t, err)
	_, err = m.GetUser(u2ID)
	assert.Equal(t, entity.ErrNotFound, err)

	u3 := newFixtureUser()
	id, _ := m.CreateUser(u3.Name, u3.Email, u3.Password, u3.Address, u3.Age)
	saved, _ := m.GetUser(id)
	_ = m.UpdateUser(saved)
	err = m.DeleteUser(id)
	assert.Equal(t, nil, err)
}
