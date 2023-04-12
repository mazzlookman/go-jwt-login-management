package service

import (
	"login-management-go/model/domain"
	"net/http"
)

type SessionService interface {
	Create(w http.ResponseWriter, userID int) domain.Session
	CurrentUser(r *http.Request) domain.User
	Destroy(w http.ResponseWriter, r *http.Request)
}
