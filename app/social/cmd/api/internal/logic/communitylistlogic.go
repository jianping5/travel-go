package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
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
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	userId := req.UserId
	// 若传入的 userId 为 0，则表示为当前用户
	if userId == 0 {
		userId = ctxdata.GetUidFromCtx(l.ctx)
	}
	var communityIds []int64
	if userId != -1 {
		// 获取该用户下的社区 ID 列表
		var userCommunityList []model.UserCommunity
		l.svcCtx.DB.Select("community_id").Where("user_id = ?", userId).Find(&userCommunityList)
		if len(userCommunityList) == 0 {
			return &types.CommunityListResp{
				List: nil,
			}, nil
		}
		for _, userCommunity := range userCommunityList {
			communityIds = append(communityIds, userCommunity.CommunityId)
		}
	}

	// 获取社区详情列表
	var communityViews []types.CommunityView
	var communityList []model.Community
	// 若 userId 为 -1，则表示查询所有社区
	if userId == -1 {
		l.svcCtx.DB.Model(&model.Community{}).Scan(&communityList)
	} else {
		l.svcCtx.DB.Where("id IN (?)", communityIds).Find(&communityList)
	}

	for _, community := range communityList {
		var communityView types.CommunityView
		_ = copier.Copy(&communityView, community)
		communityView.CreateTime = tool.TimeToString(community.CreateTime)

		// 获取是否已经加入
		var id int64
		if l.svcCtx.DB.Model(&model.UserCommunity{}).Select("id").Where("user_id = ? and community_id = ?", loginUserId, community.Id).Scan(&id); id == 0 {
			communityView.IsJoined = false
		} else {
			communityView.IsJoined = true
		}

		communityViews = append(communityViews, communityView)
	}

	return &types.CommunityListResp{
		List: communityViews,
	}, nil
}
