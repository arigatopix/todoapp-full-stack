package utils

import (
	"server/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

var env = config.LoadENV()
var jwtSecret []byte = []byte(env.JWT_SECRET)

func GenerateToken(userId int) (string, error) {
	timeHour, _ := time.ParseDuration(env.JWT_EXPIRE_IN)

	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeHour).Unix(),
			Issuer:    "todoapp",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
