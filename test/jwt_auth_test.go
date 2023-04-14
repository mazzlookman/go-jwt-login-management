package test

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	authJwt "login-management-go/auth"
	"login-management-go/model/web"
	"testing"
)

var auth = authJwt.NewJWTAuth()

func TestGenerateToken(t *testing.T) {
	token, err := auth.GenerateToken(web.JWTClaims{
		SID: uuid.New().String(),
	})
	if err != nil {
		log.Println(err)
	}

	log.Println(token)
}

func TestValidateToken(t *testing.T) {
	token, err := auth.GenerateToken(web.JWTClaims{
		SID: uuid.New().String(),
	})
	if err != nil {
		log.Println(err)
	}

	validateToken, err := auth.ValidateToken(token)
	if err != nil {
		log.Println(err)
	}
	claims := validateToken.Claims.(jwt.MapClaims)
	log.Println(claims["s_id"])
	assert.Equal(t, 1, int(claims["user_id"].(float64)))
}
