package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"
	"travel/common/tool"
	"travel/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityDynamicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityDynamicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityDynamicListLogic {
	return &CommunityDynamicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityDynamicListLogic) CommunityDynamicList(req *types.CommunityDynamicListReq) (resp *types.CommunityDynamicListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	offset := (req.PageNum - 1) * req.PageSize
	var dynamics []model.Dynamic
	var total int64
	tx := l.svcCtx.DB.Model(&model.Dynamic{})

	// Type: 最热动态	最新发布	最近常看	已加入社区
	dynamicType := req.Type
	switch enum.DynamicType(dynamicType) {
	case enum.Hot:
		tx = tx.Order("like_count DESC")
	case enum.Latest:
		tx = tx.Order("create_time DESC")
	case enum.Recent:
		// todo: 最近常看，需要结合数据模块

	default:
		return nil, errors.Wrap(xerr.NewErrMsg("参数错误"), "参数不匹配")
	}

	// 仅查看已加入社区的
	var communityIds []int
	if req.JoinedSwitch {
		l.svcCtx.DB.Model(&model.UserCommunity{}).Where("user_id = ?", loginUserId).Pluck("community_id", &communityIds)
		// 反之，若社区为空，则表示没有对应动态，直接返回，不能只是说把条件去除
		if len(communityIds) > 0 {
			tx.Where("community_id in (?)", communityIds)
		} else {
			return &types.CommunityDynamicListResp{}, nil
		}
	}

	// 分页之前，获取满足对应条件的记录总数
	countTx := tx
	countTx.Count(&total)

	// 分页
	tx.Offset(offset).Limit(req.PageSize).Find(&dynamics)

	var dynamicViews []types.CommunityDynamicView
	for _, dynamic := range dynamics {
		var dynamicView types.CommunityDynamicView
		_ = copier.Copy(&dynamicView, &dynamic)
		dynamicView.CreateTime = tool.TimeToString(dynamic.CreateTime)
		// 用户信息
		var userInfoView types.UserInfoView
		userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
			Id:          dynamic.UserId,
			LoginUserId: loginUserId,
		})
		_ = copier.Copy(&userInfoView, &userInfo)
		dynamicView.UserInfo = userInfoView
		// 社区信息
		var communityView types.CommunityView
		l.svcCtx.DB.Take(&model.Community{}).Where("id = ?", dynamic.CommunityId).Scan(&communityView)
		dynamicView.Community = communityView

		// 是否点赞
		var isLiked bool
		l.svcCtx.DB.Model(&model.Like{}).Select("liked_status").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.DYNAMIC, dynamic.Id).Scan(&isLiked)
		dynamicView.IsLiked = isLiked

		dynamicViews = append(dynamicViews, dynamicView)
	}

	return &types.CommunityDynamicListResp{
		List:  dynamicViews,
		Total: int(total),
	}, nil
}
