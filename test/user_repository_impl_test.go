package test

import (
	"github.com/stretchr/testify/assert"
	"login-management-go/app"
	"login-management-go/model/domain"
	"login-management-go/repository"
	"testing"
)

var userRepository = repository.NewUserRepository(app.DBConnection())

func TestCreateUser(t *testing.T) {

	user := domain.User{
		Name:     "Pram",
		Password: "123",
	}
	save := userRepository.Save(user)
	assert.Equal(t, 3, save.ID)
}

func TestUpdateUser(t *testing.T) {
	update := userRepository.Update(domain.User{
		ID:       3,
		Name:     "Pramuja",
		Password: "1234",
	})

	assert.Equal(t, "1234", update.Password)
}

func TestFindByIDUser(t *testing.T) {
	user := userRepository.FindByID(2)
	assert.Equal(t, "Lookman", user.Name)
}

func TestDeleteByIDUser(t *testing.T) {
	userRepository.DeleteByID(3)

	user := userRepository.FindByID(2)
	assert.Equal(t, 0, user.ID)
	assert.Equal(t, "", user.Name)
	assert.Equal(t, "", user.Password)
}
