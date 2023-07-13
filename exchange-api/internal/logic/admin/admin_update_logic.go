package admin

import (
	"context"
	"net/http"

	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdateLogic struct {
	logx.Logger
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUpdateLogic(r *http.Request, svcCtx *svc.ServiceContext) *AdminUpdateLogic {
	return &AdminUpdateLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *AdminUpdateLogic) AdminUpdate(req *types.AdminUpdateReq) (resp *types.AdminCommonResp, err error) {
	// todo: add your logic here and delete this line
	l.Infow("AdminUpdate start", logx.Field("req", req))

	return
}
