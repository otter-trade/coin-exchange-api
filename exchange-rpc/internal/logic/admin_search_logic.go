package logic

import (
	"context"

	"github.com/otter-trade/coin-exchange-api/exchange-rpc/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminSearchLogic {
	return &AdminSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 管理员列表
func (l *AdminSearchLogic) AdminSearch(in *pb.AdminSearchReq) (resp *pb.AdminSearchResp, err error) {
	// todo: add your logic here and delete this line
	l.Infow("AdminSearch start", logx.Field("in", in))
	resp = &pb.AdminSearchResp{}

	return resp, nil
}
