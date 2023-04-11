package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"login-management-go/helper"
)

func DBConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/login_management_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}