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

type UserHomeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomeListLogic {
	return &UserHomeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomeListLogic) UserHomeList(req *types.UserHomeListReq) (resp *types.UserHomeListResp, err error) {
	userId := req.UserId
	// 若传入的 userId 为 0，则表示为当前用户
	if userId == 0 {
		userId = ctxdata.GetUidFromCtx(l.ctx)
	}
	// 最新文章
	var recentArticles []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "user_id", "title", "cover_url", "like_count", "create_time").
		Where("user_id = ? and item_type = ?", userId, enum.ARTICLE).Order("create_time DESC").Limit(5).Scan(&recentArticles)
	l.SetContentInfo(&recentArticles)
	// 最新视频
	var recentVideos []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "user_id", "title", "cover_url", "like_count", "create_time").
		Where("user_id = ? and item_type = ?", userId, enum.VIDEO).Order("create_time DESC").Limit(5).Scan(&recentVideos)
	l.SetContentInfo(&recentVideos)
	// TODO: 推荐是否考虑评论的正负性
	// 文章推荐
	var recommendArticles []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "user_id", "title", "cover_url", "like_count", "create_time").
		Where("user_id = ? and item_type = ?", userId, enum.ARTICLE).Order("like_count+comment_count+favor_count DESC").Limit(5).Scan(&recommendArticles)
	l.SetContentInfo(&recommendArticles)
	// 视频推荐
	var recommendVideos []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "user_id", "title", "cover_url", "like_count", "create_time").
		Where("user_id = ? and item_type = ?", userId, enum.VIDEO).Order("like_count+comment_count+favor_count DESC").Limit(5).Scan(&recommendVideos)
	l.SetContentInfo(&recommendVideos)

	return &types.UserHomeListResp{
		RecentArticleList:    recentArticles,
		RecentVideoList:      recentVideos,
		RecommendArticleList: recommendArticles,
		RecommendVideoList:   recommendVideos,
	}, nil
}

func (l *UserHomeListLogic) SetContentInfo(contents *[]types.ContentView) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	for i, v := range *contents {
		// 用户信息
		var userInfoView types.UserInfoView
		userId := v.UserId
		info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId, LoginUserId: loginUserId})
		_ = copier.Copy(&userInfoView, &info)
		(*contents)[i].UserInfo = userInfoView
	}
}
