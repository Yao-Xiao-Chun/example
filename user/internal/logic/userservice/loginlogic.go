package userservicelogic

import (
	"context"
	"example/order/i-kun.vip/pb/order"
	"fmt"

	"example/user/i-kun.vip/pb/user"
	"example/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line
	ping, err := l.svcCtx.OrderRpc.Ping(l.ctx, &order.Request{
		Ping: "ping",
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("测试返回值", ping.Pong)

	return &user.LoginResp{}, nil
}
