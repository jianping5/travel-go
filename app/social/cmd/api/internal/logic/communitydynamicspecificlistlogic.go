package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"
	"travel/common/tool"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityDynamicSpecificListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityDynamicSpecificListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityDynamicSpecificListLogic {
	return &CommunityDynamicSpecificListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityDynamicSpecificListLogic) CommunityDynamicSpecificList(req *types.CommunityDynamicSpecificListReq) (resp *types.CommunityDynamicSpecificListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	offset := (req.PageNum - 1) * req.PageSize
	var dynamics []model.Dynamic
	var total int64
	tx := l.svcCtx.DB.Model(&model.Dynamic{}).Where("community_id = ?", req.CommunityId)

	// Type: 最热动态	最新发布	最近常看	已加入社区
	sortType := req.SortType
	switch enum.SortType(sortType) {
	case enum.Popular:
		tx = tx.Order("like_count DESC")
		break
	case enum.Newest:
		tx = tx.Order("create_time DESC")
		break
	case enum.Oldest:
		tx = tx.Order("create_time ASC")
		break
	default:
		return nil, errors.Wrap(xerr.NewErrMsg("参数错误"), "参数不匹配")
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

	return &types.CommunityDynamicSpecificListResp{
		List:  dynamicViews,
		Total: int(total),
	}, nil

	return
}
