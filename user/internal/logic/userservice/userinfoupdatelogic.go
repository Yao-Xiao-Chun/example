package userservicelogic

import (
	"context"

	"example/user/i-kun.vip/pb/user"

	"example/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(in *user.UserInfoUpdateReq) (*user.UserInfoUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoUpdateResp{}, nil
}
