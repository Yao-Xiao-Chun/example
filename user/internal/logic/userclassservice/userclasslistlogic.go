package userclassservicelogic

import (
	"context"

	"example/user/github.com/example/user"
	"example/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserClassListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserClassListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserClassListLogic {
	return &UserClassListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserClassListLogic) UserClassList(in *user.UserClassListReq) (*user.UserClassListResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserClassListResp{}, nil
}
