package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyrightDetailReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyrightDetailReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyrightDetailReqLogic {
	return &CopyrightDetailReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyrightDetailReqLogic) CopyrightDetailReq(req *types.CopyrightDetailReq) (resp *types.CopyrightDetailResp, err error) {
	// 版权信息
	var copyright types.CopyrightView
	l.svcCtx.DB.Model(&model.Copyright{}).Where("userId = ? and itemType = ? and itemId = ?", req.UserId, req.ItemType, req.ItemId).Scan(&copyright)

	// 用户信息
	var userInfoView types.UserInfoView
	userId := copyright.UserId
	userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId})
	_ = copier.Copy(&userInfoView, &userInfo)

	return &types.CopyrightDetailResp{
		Copyright: copyright,
		UserInfo:  userInfoView,
	}, nil
}
