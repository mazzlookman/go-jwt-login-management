package repository

import (
	"gorm.io/gorm"
	"login-management-go/helper"
	"login-management-go/model/domain"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r *UserRepositoryImpl) FindByEmail(email string) domain.User {
	user := domain.User{}
	err := r.db.Find(&user, "email=?", email).Error
	helper.PanicIfError(err)

	return user
}

func (r *UserRepositoryImpl) Save(user domain.User) domain.User {
	err := r.db.Create(&user).Error
	helper.PanicIfError(err)

	return user
}

func (r *UserRepositoryImpl) Update(user domain.User) domain.User {
	err := r.db.Save(&user).Error
	helper.PanicIfError(err)

	return user
}

func (r *UserRepositoryImpl) FindByID(userID int) domain.User {
	user := domain.User{}
	err := r.db.Where("id=?", userID).Find(&user).Error
	helper.PanicIfError(err)

	return user
}

func (r *UserRepositoryImpl) DeleteByID(userID int) {
	user := domain.User{}
	err := r.db.Delete(&user, userID).Error
	helper.PanicIfError(err)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
