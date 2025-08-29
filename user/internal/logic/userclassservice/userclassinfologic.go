package userclassservicelogic

import (
	"context"

	"example/user/i-kun.vip/pb/user"
	"example/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserClassInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserClassInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserClassInfoLogic {
	return &UserClassInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserClassInfoLogic) UserClassInfo(in *user.UserClassInfoReq) (*user.UserClassInfoResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserClassInfoResp{}, nil
}
