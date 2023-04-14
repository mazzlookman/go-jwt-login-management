package service

import (
	"github.com/golang-jwt/jwt/v5"
	"login-management-go/auth"
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
	auth.JWTAuth
}

func (s *SessionServiceImpl) Create(w http.ResponseWriter, token string, userID int) domain.Session {
	validateToken, err := s.JWTAuth.ValidateToken(token)
	helper.PanicIfError(err)

	claims := validateToken.Claims.(jwt.MapClaims)
	sess := domain.Session{
		ID:     claims["s_id"].(string),
		UserID: userID,
	}

	session := s.SessionRepository.Save(sess)

	cookie := new(http.Cookie)
	cookie.Name = cookie_name
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(1 * time.Hour)
	http.SetCookie(w, cookie)

	return session
}

func (s *SessionServiceImpl) CurrentUser(r *http.Request) domain.User {
	cookie, err := r.Cookie(cookie_name)
	helper.PanicIfError(err)

	validateToken, err := s.JWTAuth.ValidateToken(cookie.Value)
	helper.PanicIfError(err)

	claims := validateToken.Claims.(jwt.MapClaims)
	session := s.SessionRepository.FindByID(claims["s_id"].(string))
	user := s.UserRepository.FindByID(session.UserID)

	return user
}

func (s *SessionServiceImpl) Destroy(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(cookie_name)
	helper.PanicIfError(err)

	validateToken, err := s.JWTAuth.ValidateToken(c.Value)
	helper.PanicIfError(err)

	claims := validateToken.Claims.(jwt.MapClaims)
	s.SessionRepository.DeleteByID(claims["s_id"].(string))

	c.Expires = time.Unix(0, 0)
	c.MaxAge = -1

	http.SetCookie(w, c)
}

func NewSessionService(userRepository repository.UserRepository, sessionRepository repository.SessionRepository, jwtAuth auth.JWTAuth) SessionService {
	return &SessionServiceImpl{UserRepository: userRepository, SessionRepository: sessionRepository, JWTAuth: jwtAuth}
}
