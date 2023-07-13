package svc

import (
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/config"
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/middleware"
	"github.com/otter-trade/coin-exchange-api/exchange-rpc/exchange"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	CheckSignMiddleware rest.Middleware
	TokenMiddleware     rest.Middleware
	CorsMiddleware      rest.Middleware
	Exchange            exchange.Exchange
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		Exchange:            exchange.NewExchange(zrpc.MustNewClient(c.ExchangeRpc)),
		CheckSignMiddleware: middleware.NewCheckSignMiddleware().Handle,
		TokenMiddleware:     middleware.NewTokenMiddleware().Handle,
		CorsMiddleware:      middleware.NewCorsMiddleware().Handle,
	}
}
