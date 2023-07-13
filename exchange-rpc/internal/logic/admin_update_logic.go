package logic

import (
	"context"

	"github.com/otter-trade/coin-exchange-api/exchange-rpc/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUpdateLogic {
	return &AdminUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新管理员
func (l *AdminUpdateLogic) AdminUpdate(in *pb.AdminUpdateReq) (resp *pb.AdminUpdateResp, err error) {
	// todo: add your logic here and delete this line
	l.Infow("AdminUpdate start", logx.Field("in", in))
	resp = &pb.AdminUpdateResp{}

	return resp, nil
}
