package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/rpc/social"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyrightDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyrightDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyrightDetailLogic {
	return &CopyrightDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyrightDetailLogic) CopyrightDetail(req *types.CopyrightDetailReq) (resp *types.CopyrightDetailResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	// 版权信息
	var copyright types.CopyrightView
	detail, err := l.svcCtx.SocialRpc.CopyrightDetail(l.ctx, &social.CopyrightDetailReq{Id: req.Id})
	_ = copier.Copy(&copyright, &detail)

	// 用户信息
	var userInfoView types.UserInfoView
	userId := copyright.UserId
	userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId, LoginUserId: loginUserId})
	_ = copier.Copy(&userInfoView, &userInfo)

	return &types.CopyrightDetailResp{
		Copyright: copyright,
		UserInfo:  userInfoView,
	}, nil
}
