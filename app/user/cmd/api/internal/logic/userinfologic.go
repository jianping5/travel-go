package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"travel/app/user/cmd/api/internal/svc"
	"travel/app/user/cmd/api/internal/types"
	"travel/app/user/cmd/rpc/user"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userId := req.Id

	// 从当前登录态中获取
	// userId = ctxdata.GetUidFromCtx(l.ctx)

	userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UserInfoResp{}
	_ = copier.Copy(resp, userInfoResp)

	return resp, nil
}
