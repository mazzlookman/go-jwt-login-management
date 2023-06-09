package repository

import "login-management-go/model/domain"

type SessionRepository interface {
	Save(session domain.Session) domain.Session
	FindByID(sessionID string) domain.Session
	DeleteByID(sessionID string)
}
