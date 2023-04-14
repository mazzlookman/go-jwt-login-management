package main

import (
	"login-management-go/app"
)

func main() {
	router := app.NewRouter()
	router.Run(":8080")
}
