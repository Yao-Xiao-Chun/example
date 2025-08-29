package orderservicelogic

import (
	"context"
	"example/order/i-kun.vip/pb/order"
	"example/order/internal/svc"
	"github.com/google/uuid"

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
	//生成一个订单号
	orderNo := "example-" + uuid.New().String() //"uuid.New().String()
	if in.OrderNo == "" {
		in.OrderNo = orderNo
	}

	//创建订单
	//验证商品充足
	//扣减库存
	//支付
	//完成

	return &order.OrderCreateResp{
		OrderNo: in.OrderNo,
		PayNo:   "payNo-2025-08-27",
	}, nil
}
