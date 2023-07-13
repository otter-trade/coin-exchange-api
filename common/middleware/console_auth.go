package middleware

import (
	"bytes"
	"github.com/gookit/filter"
	"github.com/otter-trade/coin-exchange-api/common/i18n"
	"github.com/otter-trade/coin-exchange-api/common/xresp"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type ConsoleAuthMiddleware struct {
	RedisClient *redis.Redis
}

func NewConsoleAuthMiddleware(redisClient *redis.Redis) *AuthMiddleware {
	return &AuthMiddleware{
		RedisClient: redisClient,
	}
}

func (m *ConsoleAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("example middle")
	return func(w http.ResponseWriter, r *http.Request) {
		userId, _ := filter.String(r.Context().Value("userId"))

		var buf bytes.Buffer
		buf.WriteString("c:")
		buf.WriteString(userId)

		_, err := m.RedisClient.Get(buf.String())
		if err != nil {
			xresp.Fail(r, w, i18n.NewApiError(i18n.Unauthorized))
			return
		}

		next(w, r)
	}
}
