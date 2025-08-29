package orderservicelogic

import (
	"context"

	"example/order/i-kun.vip/pb/order"
	"example/order/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderRevertLogic {
	return &OrderRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderRevertLogic) OrderRevert(in *order.OrderRollbackReq) (*order.OrderRollbackResp, error) {
	// todo: add your logic here and delete this line

	return &order.OrderRollbackResp{}, nil
}
