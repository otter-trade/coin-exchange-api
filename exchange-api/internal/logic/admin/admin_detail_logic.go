package admin

import (
	"context"
	"net/http"

	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDetailLogic struct {
	logx.Logger
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminDetailLogic(r *http.Request, svcCtx *svc.ServiceContext) *AdminDetailLogic {
	return &AdminDetailLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *AdminDetailLogic) AdminDetail(req *types.AdminDetailReq) (resp *types.AdminDetailResp, err error) {
	// todo: add your logic here and delete this line
	l.Infow("AdminDetail start", logx.Field("req", req))

	return
}
