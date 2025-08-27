package userclassservicelogic

import (
	"context"

	"example/user/github.com/example/user"
	"example/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserClassUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserClassUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserClassUpdateLogic {
	return &UserClassUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserClassUpdateLogic) UserClassUpdate(in *user.UserClassUpdateReq) (*user.UserClassUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserClassUpdateResp{}, nil
}
