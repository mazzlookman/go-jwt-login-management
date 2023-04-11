package repository

import "login-management-go/model/domain"

type SessionRepository interface {
	Save(session domain.Session) domain.Session
	FindByID(sessionID int) domain.Session
	DeleteByID(sessionID int)
}
