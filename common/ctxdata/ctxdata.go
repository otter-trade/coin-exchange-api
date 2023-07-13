package ctxdata

import (
	"errors"
	"github.com/otter-trade/coin-exchange-api/common/helpers"
	"net/http"
	"strings"
)

func GetUserId(r *http.Request, accessSecret string) (int64, error) {
	claims, err := ParseToken(r, accessSecret)
	if claims != nil {
		return claims.UserId, err
	}
	return 0, err
}

func GetUuid(r *http.Request, accessSecret string) (string, error) {
	claims, err := ParseToken(r, accessSecret)
	if claims != nil {
		return claims.Uuid, err
	}
	return "", err
}

func ParseToken(r *http.Request, accessSecret string) (*helpers.Claims, error) {
	authorization := r.Header.Get("Authorization")
	authorizationRes := strings.Split(authorization, " ")
	if len(authorizationRes) > 1 {
		claims, err := helpers.ParseToken(authorizationRes[1], accessSecret)
		if err != nil {
			return nil, err
		}
		return claims, nil
	}
	return nil, errors.New("Invalid authorization")
}
