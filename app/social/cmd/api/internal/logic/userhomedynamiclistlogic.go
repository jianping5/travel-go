package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomeDynamicListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomeDynamicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomeDynamicListLogic {
	return &UserHomeDynamicListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomeDynamicListLogic) UserHomeDynamicList(req *types.UserHomeDynamicListReq) (resp *types.UserHomeDynamicListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	userId := req.UserId
	offset := (req.PageNum - 1) * req.PageSize
	var total int64
	var dynamics []types.CommunityDynamicView
	tx := l.svcCtx.DB.Model(&model.Dynamic{}).
		Where("user_id = ?", userId).Order("create_time DESC")

	// 记录总数
	countTx := tx
	countTx.Count(&total)

	tx.Offset(offset).Limit(req.PageSize).Scan(&dynamics)

	for i, d := range dynamics {
		// 用户信息
		var userInfoView types.UserInfoView
		userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{
			Id:          userId,
			LoginUserId: loginUserId,
		})
		_ = copier.Copy(&userInfoView, &userInfo)
		dynamics[i].UserInfo = userInfoView
		// 社区信息
		var communityView types.CommunityView
		l.svcCtx.DB.Take(&model.Community{}).Where("id = ?", d.CommunityId).Scan(&communityView)
		dynamics[i].Community = communityView

		// 未登录
		if loginUserId == 0 {
			dynamics[i].IsLiked = false
		} else {
			// 是否点赞
			var isLiked bool
			l.svcCtx.DB.Model(&model.Like{}).Select("liked_status").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.DYNAMIC, d.Id).Scan(&isLiked)
			dynamics[i].IsLiked = isLiked
		}
	}

	return &types.UserHomeDynamicListResp{
		List:  dynamics,
		Total: int(total),
	}, nil
}
