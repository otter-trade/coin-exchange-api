package logic

import (
	"context"

	"github.com/otter-trade/coin-exchange-api/exchange-rpc/internal/svc"
	"github.com/otter-trade/coin-exchange-api/exchange-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminDetailsLogic {
	return &AdminDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取管理员
func (l *AdminDetailsLogic) AdminDetails(in *pb.AdminDetailsReq) (resp *pb.AdminDetailsResp, err error) {
	// todo: add your logic here and delete this line
	l.Infow("AdminDetails start", logx.Field("in", in))
	resp = &pb.AdminDetailsResp{}

	return resp, nil
}
