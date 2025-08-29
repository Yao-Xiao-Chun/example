package userroleservicelogic

import (
	"context"

	"example/user/i-kun.vip/pb/user"
	"example/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleInfoLogic {
	return &UserRoleInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleInfoLogic) UserRoleInfo(in *user.UserRoleInfoReq) (*user.UserRoleInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserRoleInfoResp{}, nil
}
