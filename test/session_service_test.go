package test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"login-management-go/app"
	"login-management-go/model/web"
	"login-management-go/repository"
	"login-management-go/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var sessionService = service.NewSessionService(
	repository.NewUserRepository(app.DBConnection()),
	repository.NewSessionRepository(app.DBConnection()),
)

func TestCreate(t *testing.T) {
	user, _ := userService.Login(web.UserLoginInput{
		Email:    "mario@test.com",
		Password: "123",
	})

	w := httptest.NewRecorder()
	session := sessionService.Create(w, user.ID)

	for _, cookie := range w.Result().Cookies() {
		assert.Equal(t, cookie.Name, "I-BELL-A")
		assert.Equal(t, cookie.Value, session.ID)
	}
}

func TestCurrentUser(t *testing.T) {
	//user login
	user, _ := userService.Login(web.UserLoginInput{
		Email:    "ucup@test.com",
		Password: "123",
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	//generate cookie
	session := sessionService.Create(w, user.ID)
	for _, cookie := range w.Result().Cookies() {
		assert.Equal(t, cookie.Name, "I-BELL-A")
		assert.Equal(t, cookie.Value, session.ID)
		r.AddCookie(cookie)
	}

	//get user by cookie
	currentUser := sessionService.CurrentUser(r)
	assert.Equal(t, "ucup@test.com", currentUser.Email)
	log.Println(currentUser)
}

func TestDestroy(t *testing.T) {
	//user login
	user, _ := userService.Login(web.UserLoginInput{
		Email:    "mario@test.com",
		Password: "123",
	})

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	//generate cookie
	session := sessionService.Create(w, user.ID)
	for _, cookie := range w.Result().Cookies() {
		assert.Equal(t, cookie.Name, "i-bell-a")
		assert.Equal(t, cookie.Value, session.ID)
		r.AddCookie(cookie)
	}
	log.Println(session)

	//destroy cookie
	sessionService.Destroy(w, r)
	cookieDelete, _ := r.Cookie("i-bell-a")
	log.Println("After Delete")
	log.Println(cookieDelete)

	//get current user
	currentUser := sessionService.CurrentUser(r)
	log.Println("User After Delete")
	log.Println(currentUser)
}
