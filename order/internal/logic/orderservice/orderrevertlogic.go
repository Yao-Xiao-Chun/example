package orderservicelogic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	if in.OrderNo == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument")
	}

	return &order.OrderRollbackResp{}, nil
}
