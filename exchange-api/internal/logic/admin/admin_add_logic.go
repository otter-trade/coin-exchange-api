package admin

import (
	"context"
	"net/http"

	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminAddLogic struct {
	logx.Logger
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminAddLogic(r *http.Request, svcCtx *svc.ServiceContext) *AdminAddLogic {
	return &AdminAddLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *AdminAddLogic) AdminAdd(req *types.AdminAddReq) (resp *types.AdminCommonResp, err error) {
	// todo: add your logic here and delete this line
	l.Infow("AdminAdd start", logx.Field("req", req))

	return
}
