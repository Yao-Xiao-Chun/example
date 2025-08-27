package orderservicelogic

import (
	"context"

	"example/order/i-kun.vip/pb/order"
	"example/order/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderCreateLogic) OrderCreate(in *order.OrderCreateReq) (*order.OrderCreateResp, error) {
	// todo: add your logic here and delete this line

	return &order.OrderCreateResp{
		OrderNo: in.OrderNo,
		PayNo:   "payNo-2025-08-27",
	}, nil
}
