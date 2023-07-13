package helpers

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Exp         int64    `json:"exp"`
	Iat         int64    `json:"iat"`
	UserId      int64    `json:"userId"`
	TokenTime   int64    `json:"tokenTime"`
	Uuid        string   `json:"uuid"`
	NickName    string   `json:"nickName"`
	UserName    string   `json:"userName"`
	IsAdmin     bool     `json:"isAdmin"`
	RoleKey     string   `json:"roleKey"`
	RoleId      int64    `json:"roleId"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

func ParseToken(token string, jwtSecret string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, errors.New("couldn't handle this token")
}
