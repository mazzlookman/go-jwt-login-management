package repository

import "login-management-go/model/domain"

type UserRepository interface {
	Save(user domain.User) domain.User
	Update(user domain.User) domain.User
	FindByID(userID int) domain.User
	FindByEmail(email string) domain.User
	DeleteByID(userID int)
}
