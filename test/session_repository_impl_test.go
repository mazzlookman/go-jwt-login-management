package test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"login-management-go/app"
	"login-management-go/model/domain"
	"login-management-go/repository"
	"testing"
)

var sessionRepo = repository.NewSessionRepository(app.DBConnection())

func TestCreateSession(t *testing.T) {
	session := sessionRepo.Save(domain.Session{
		ID:     uuid.New().String(),
		UserID: 1,
	})
	assert.Equal(t, 1, session.UserID)
}

func TestFindByIDSession(t *testing.T) {
	session := sessionRepo.FindByID("fdd249d2-3ee7-4438-866e-ecc96444556b")
	assert.Equal(t, 2, session.UserID)
}

func TestDeleteByIDSession(t *testing.T) {
	sessionRepo.DeleteByID("fdd249d2-3ee7-4438-866e-ecc96444556b")

	session := sessionRepo.FindByID("fdd249d2-3ee7-4438-866e-ecc96444556b")
	assert.Equal(t, 0, session.UserID)
}
