package orderpaylogic

import (
	"context"

	"example/pay/i-kun.vip/pb/pay"
	"example/pay/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPayLogic {
	return &OrderPayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderPayLogic) OrderPay(in *pay.OrderPayReq) (*pay.OrderPayResp, error) {
	// todo: add your logic here and delete this line

	return &pay.OrderPayResp{}, nil
}
