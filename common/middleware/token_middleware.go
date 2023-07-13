package middleware

import (
	"github.com/otter-trade/coin-exchange-api/common/ctxdata"
	"github.com/otter-trade/coin-exchange-api/common/i18n"
	"github.com/otter-trade/coin-exchange-api/common/xkeys"
	"github.com/otter-trade/coin-exchange-api/common/xresp"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type TokenMiddleware struct {
	AccessSecret string `json:"accessSecret"`
	AccessExpire int    `json:"accessExpire"`
	RedisClient  *redis.Redis
}

func NewTokenMiddleware(accessSecret string, accessExpire int, redisClient *redis.Redis) *TokenMiddleware {
	return &TokenMiddleware{AccessSecret: accessSecret, AccessExpire: accessExpire, RedisClient: redisClient}
}

func (t *TokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := ctxdata.ParseToken(r, t.AccessSecret)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			xresp.Fail(r, w, i18n.NewApiError(i18n.Unauthorized))
			return
		}

		if claims.Exp < time.Now().Unix() {
			w.WriteHeader(http.StatusUnauthorized)
			xresp.Fail(r, w, i18n.NewApiError(i18n.Unauthorized))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			xresp.Fail(r, w, i18n.NewApiError(i18n.Unauthorized))
			return
		}

		// 1 小时不操作则过期
		consoleTokenKey := xkeys.LoginTokenKey(claims.UserId)
		tokenTime, err := t.RedisClient.Get(consoleTokenKey)
		if err != nil || tokenTime == "" {
			w.WriteHeader(http.StatusUnauthorized)
			xresp.Fail(r, w, status.Error(i18n.Unauthorized, err.Error()))

			return
		}

		// 1 小时不操作则过期
		err = t.RedisClient.Expire(consoleTokenKey, t.AccessExpire)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			xresp.Fail(r, w, status.Error(i18n.Unauthorized, err.Error()))
			return
		}

		next(w, r)
	}
}
