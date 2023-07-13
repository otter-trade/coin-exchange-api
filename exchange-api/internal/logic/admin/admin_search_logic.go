package admin

import (
	"context"
	"net/http"

	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminSearchLogic struct {
	logx.Logger
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminSearchLogic(r *http.Request, svcCtx *svc.ServiceContext) *AdminSearchLogic {
	return &AdminSearchLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *AdminSearchLogic) AdminSearch(req *types.AdminSearchReq) (resp *types.AdminSearchResp, err error) {
	// todo: add your logic here and delete this line
	l.Infow("AdminSearch start", logx.Field("req", req))

	return
}
