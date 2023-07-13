package admin

import (
	"context"
	"net/http"

	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDeleteLogic struct {
	logx.Logger
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminDeleteLogic(r *http.Request, svcCtx *svc.ServiceContext) *AdminDeleteLogic {
	return &AdminDeleteLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *AdminDeleteLogic) AdminDelete(req *types.AdminDeleteReq) (resp *types.AdminCommonResp, err error) {
	// todo: add your logic here and delete this line
	l.Infow("AdminDelete start", logx.Field("req", req))

	return
}
