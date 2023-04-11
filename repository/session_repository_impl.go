package repository

import (
	"gorm.io/gorm"
	"login-management-go/helper"
	"login-management-go/model/domain"
)

type SessionRepositoryImpl struct {
	db *gorm.DB
}

func (r *SessionRepositoryImpl) Save(session domain.Session) domain.Session {
	err := r.db.Create(&session).Error
	helper.PanicIfError(err)

	return session
}

func (r *SessionRepositoryImpl) FindByID(sessionID string) domain.Session {
	session := domain.Session{}
	err := r.db.Find(&session, "id=?", sessionID).Error
	helper.PanicIfError(err)

	return session
}

func (r *SessionRepositoryImpl) DeleteByID(sessionID string) {
	err := r.db.Delete(&domain.Session{}, "id=?", sessionID).Error
	helper.PanicIfError(err)
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &SessionRepositoryImpl{db: db}
}
