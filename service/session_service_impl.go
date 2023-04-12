package service

import (
	"github.com/google/uuid"
	"login-management-go/helper"
	"login-management-go/model/domain"
	"login-management-go/repository"
	"net/http"
	"time"
)

var cookie_name = "i-bell-a"

type SessionServiceImpl struct {
	repository.UserRepository
	repository.SessionRepository
}

func (s *SessionServiceImpl) Create(w http.ResponseWriter, userID int) domain.Session {
	sess := domain.Session{
		ID:     uuid.New().String(),
		UserID: userID,
	}

	session := s.SessionRepository.Save(sess)

	cookie := new(http.Cookie)
	cookie.Name = cookie_name
	cookie.Value = session.ID
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(1 * time.Hour)
	http.SetCookie(w, cookie)

	return session
}

func (s *SessionServiceImpl) CurrentUser(r *http.Request) domain.User {
	cookie, err := r.Cookie(cookie_name)
	helper.PanicIfError(err)

	session := s.SessionRepository.FindByID(cookie.Value)
	user := s.UserRepository.FindByID(session.UserID)

	return user
}

func (s *SessionServiceImpl) Destroy(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(cookie_name)
	helper.PanicIfError(err)
	s.SessionRepository.DeleteByID(c.Value)

	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1

	http.SetCookie(w, c)
}

func NewSessionService(userRepository repository.UserRepository, sessionRepository repository.SessionRepository) SessionService {
	return &SessionServiceImpl{UserRepository: userRepository, SessionRepository: sessionRepository}
}
