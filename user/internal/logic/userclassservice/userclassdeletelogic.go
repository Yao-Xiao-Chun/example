package userclassservicelogic

import (
	"context"

	"example/user/github.com/example/user"
	"example/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserClassDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserClassDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserClassDeleteLogic {
	return &UserClassDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserClassDeleteLogic) UserClassDelete(in *user.UserClassDeleteReq) (*user.UserClassDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserClassDeleteResp{}, nil
}
