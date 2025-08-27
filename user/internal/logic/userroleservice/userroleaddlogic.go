package userroleservicelogic

import (
	"context"

	"example/user/github.com/example/user"
	"example/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleAddLogic {
	return &UserRoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleAddLogic) UserRoleAdd(in *user.UserRoleAddReq) (*user.UserRoleAddResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserRoleAddResp{}, nil
}
