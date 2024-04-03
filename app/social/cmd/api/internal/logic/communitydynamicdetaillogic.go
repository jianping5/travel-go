package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"
	"travel/common/tool"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityDynamicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityDynamicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityDynamicDetailLogic {
	return &CommunityDynamicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityDynamicDetailLogic) CommunityDynamicDetail(req *types.CommunityDynamicDetailReq) (resp *types.CommunityDynamicDetailResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var dynamic model.Dynamic
	l.svcCtx.DB.Model(&model.Dynamic{}).Where("id = ?", req.Id).Find(&dynamic)

	var dynamicView types.CommunityDynamicView
	_ = copier.Copy(&dynamicView, &dynamic)
	dynamicView.CreateTime = tool.TimeToString(dynamic.CreateTime)

	// 用户信息
	var userInfoView types.UserInfoView
	userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
		Id: dynamic.UserId,
	})
	_ = copier.Copy(&userInfoView, &userInfo)
	dynamicView.UserInfo = userInfoView

	// 社区信息
	var communityView types.CommunityView
	l.svcCtx.DB.Take(&model.Community{}).Where("id = ?", dynamic.CommunityId).Scan(&communityView)
	dynamicView.Community = communityView

	// 是否点赞
	if loginUserId == 0 {
		dynamicView.IsLiked = false
	} else {
		var isLiked bool
		l.svcCtx.DB.Model(&model.Like{}).Select("liked_status").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.DYNAMIC, dynamic.Id).Scan(&isLiked)
		dynamicView.IsLiked = isLiked
	}

	return &types.CommunityDynamicDetailResp{
		DynamicDetail: dynamicView,
	}, nil
}
