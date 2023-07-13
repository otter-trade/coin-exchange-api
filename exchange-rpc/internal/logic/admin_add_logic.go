package logic

import (
	"context"

	"github.com/otter-trade/coin-exchange-api/exchange-rpc/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminAddLogic {
	return &AdminAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------管理员-----------------------
func (l *AdminAddLogic) AdminAdd(in *pb.AdminAddReq) (resp *pb.AdminAddResp, err error) {
	// todo: add your logic here and delete this line
	l.Infow("AdminAdd start", logx.Field("in", in))
	resp = &pb.AdminAddResp{}

	return resp, nil
}
