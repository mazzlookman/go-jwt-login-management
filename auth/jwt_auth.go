package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"login-management-go/model/web"
)

var SECRET_KEY = []byte("inikeyyangsangatsupersyekali")

type JWTAuth interface {
	GenerateToken(payload web.JWTClaims) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTAuthImpl struct {
}

func NewJWTAuth() JWTAuth {
	return &JWTAuthImpl{}
}

func (j *JWTAuthImpl) GenerateToken(payload web.JWTClaims) (string, error) {
	claims := jwt.MapClaims{}
	claims["s_id"] = payload.SID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (j *JWTAuthImpl) ValidateToken(token string) (*jwt.Token, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return SECRET_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	return parseToken, nil
}
