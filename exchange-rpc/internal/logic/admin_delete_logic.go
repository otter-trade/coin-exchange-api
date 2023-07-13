package logic

import (
	"context"

	"github.com/otter-trade/coin-exchange-api/exchange-rpc/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminDeleteLogic {
	return &AdminDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除管理员
func (l *AdminDeleteLogic) AdminDelete(in *pb.AdminDeleteReq) (resp *pb.AdminDeleteResp, err error) {
	// todo: add your logic here and delete this line
	l.Infow("AdminDelete start", logx.Field("in", in))
	resp = &pb.AdminDeleteResp{}

	return resp, nil
}
