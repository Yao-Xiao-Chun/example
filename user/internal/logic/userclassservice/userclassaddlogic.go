package userclassservicelogic

import (
	"context"

	"example/user/i-kun.vip/pb/user"
	"example/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserClassAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserClassAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserClassAddLogic {
	return &UserClassAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserClassAddLogic) UserClassAdd(in *user.UserClassAddReq) (*user.UserClassAddResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserClassAddResp{}, nil
}
