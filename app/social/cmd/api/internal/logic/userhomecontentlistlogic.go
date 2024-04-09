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

type UserHomeContentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomeContentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomeContentListLogic {
	return &UserHomeContentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomeContentListLogic) UserHomeContentList(req *types.UserHomeContentListReq) (resp *types.UserHomeContentListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	userId := req.UserId
	// 若传入的 userId 为 0，则表示为当前用户
	if userId == 0 {
		userId = ctxdata.GetUidFromCtx(l.ctx)
	}
	offset := (req.PageNum - 1) * req.PageSize
	var contents []types.ContentView
	var total int64
	switch enum.ItemType(req.ItemType) {
	case enum.ARTICLE:
		l.svcCtx.DB.Model(&model.Content{}).Where("user_id = ? and item_type = ?", userId, enum.ARTICLE).Count(&total)
		contents, _ = l.getSortedArticleList(req.SortType, offset, req.PageSize, userId)
		for i, a := range contents {
			// 用户信息
			var userInfoView types.UserInfoView
			userId := a.UserId
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId, LoginUserId: loginUserId})
			_ = copier.Copy(&userInfoView, &info)
			contents[i].UserInfo = userInfoView
		}
		break
	case enum.VIDEO:
		l.svcCtx.DB.Model(&model.Content{}).Where("user_id = ? and item_type = ?", userId, enum.VIDEO).Count(&total)
		contents, _ = l.getSortedVideoList(req.SortType, offset, req.PageSize, userId)
		for i, v := range contents {
			// 用户信息
			var userInfoView types.UserInfoView
			userId := v.UserId
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId, LoginUserId: loginUserId})
			_ = copier.Copy(&userInfoView, &info)
			contents[i].UserInfo = userInfoView
		}
		break
	default:
		break
	}

	return &types.UserHomeContentListResp{
		List:  contents,
		Total: int(total),
	}, nil
}

func (l *UserHomeContentListLogic) getSortedArticleList(sortType, offset, pageSize int, userId int64) ([]types.ContentView, error) {
	var contents []types.ContentView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("create_time DESC").
			Where("user_id = ? and item_type = ?", userId, enum.ARTICLE).Scan(&contents)
		break
	case enum.Popular:
		// TODO: 考虑评论的正负性
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("like_count+comment_count+favor_count DESC").
			Where("user_id = ? and item_type = ?", userId, enum.ARTICLE).Scan(&contents)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("create_time ASC").
			Where("user_id = ? and item_type = ?", userId, enum.ARTICLE).Scan(&contents)
		break
	default:
		break
	}
	return contents, nil
}

func (l *UserHomeContentListLogic) getSortedVideoList(sortType, offset, pageSize int, userId int64) ([]types.ContentView, error) {
	var contents []types.ContentView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("create_time DESC").
			Where("user_id = ? and item_type = ?", userId, enum.VIDEO).Scan(&contents)
		break
	case enum.Popular:
		// TODO: 考虑评论的正负性
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("like_count+comment_count+favor_count DESC").
			Where("user_id = ? and item_type = ?", userId, enum.VIDEO).Scan(&contents)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("create_time ASC").
			Where("user_id = ? and item_type = ?", userId, enum.VIDEO).Scan(&contents)
		break
	default:
		break
	}
	return contents, nil
}
