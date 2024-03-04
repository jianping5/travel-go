package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/social/cmd/model"
	"travel/common/tool"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityListLogic {
	return &CommunityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityListLogic) CommunityList(req *types.CommunityListReq) (resp *types.CommunityListResp, err error) {
	// 获取该用户下的社区 ID 列表
	var userCommunityList []model.UserCommunity
	l.svcCtx.DB.Select("communityId").Where("userId = ?", req.UserId).Find(&userCommunityList)
	if len(userCommunityList) == 0 {
		return nil, nil
	}
	var communityIds []int64
	for _, userCommunity := range userCommunityList {
		communityIds = append(communityIds, userCommunity.CommunityId)
	}

	// 获取社区详情列表
	var communityViews []types.CommunityView
	var communityList []model.Community
	l.svcCtx.DB.Where("id IN (?)", communityIds).Find(&communityList)
	for _, community := range communityList {
		var communityView types.CommunityView
		_ = copier.Copy(&communityView, community)
		communityView.CreateTime = tool.TimeToString(community.CreateTime)
		communityViews = append(communityViews, communityView)
	}

	return &types.CommunityListResp{
		List: communityViews,
	}, nil
}
