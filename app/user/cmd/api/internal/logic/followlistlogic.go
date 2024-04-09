package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/user/cmd/model"
	"travel/common/ctxdata"

	"travel/app/user/cmd/api/internal/svc"
	"travel/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowListLogic) FollowList(req *types.FollowListReq) (resp *types.FollowListView, err error) {
	var follows []model.Follow
	userId := req.Id
	// 若传入的 userId 为 0，则表示为当前用户
	if userId == 0 {
		userId = ctxdata.GetUidFromCtx(l.ctx)
	}
	offset := (req.PageNum - 1) * req.PageSize
	// 分页获取当前用户关注列表
	if err := l.svcCtx.DB.Where("user_id = ?", userId).Offset(int(offset)).Limit(int(req.PageSize)).Find(&follows).Error; err != nil {
		return nil, err
	}
	// 未关注
	if len(follows) == 0 {
		return nil, nil
	}
	// 获取用户信息列表
	var users []model.User
	userIds := make([]int64, len(follows))
	for k, v := range follows {
		userIds[k] = v.FollowUserId
	}
	if err := l.svcCtx.DB.Where("id IN (?)", userIds).Find(&users).Error; err != nil {
		return nil, err
	}

	var userInfos []types.UserInfoResp
	// 时间字段的转换
	for _, user := range users {
		var userInfo types.UserInfoResp
		_ = copier.Copy(&userInfo, &user)
		userInfo.CreateTime = user.CreateTime.Format("2006-01-02 15:04:05")
		userInfo.UpdateTime = user.UpdateTime.Format("2006-01-02 15:04:05")

		// 注入是否关注
		var id int64
		if l.svcCtx.DB.Model(&model.Follow{}).Select("id").Where("user_id = ? and follow_user_id = ?", userId, user.Id).Scan(&id); id == 0 {
			userInfo.IsFollowed = false
		} else {
			userInfo.IsFollowed = true
		}

		userInfos = append(userInfos, userInfo)
	}

	return &types.FollowListView{
		UserInfo: userInfos,
		Total:    int64(len(userInfos)),
	}, nil
}
