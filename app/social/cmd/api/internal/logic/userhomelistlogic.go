package logic

import (
	"context"
	"travel/app/social/cmd/model"

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
	// 最新文章
	var recentArticles []types.ContentSimpleView
	l.svcCtx.DB.Model(&model.Article{}).Select("id", "title", "coverUrl", "likeCount", "createTime").
		Where("userId = ?", userId).Order("createTime DESC").Limit(5).Scan(&recentArticles)
	// 最新视频
	var recentVideos []types.ContentSimpleView
	l.svcCtx.DB.Model(&model.Video{}).Select("id", "title", "coverUrl", "likeCount", "createTime").
		Where("userId = ?", userId).Order("createTime DESC").Limit(5).Scan(&recentVideos)
	// TODO: 推荐是否考虑评论的正负性
	// 文章推荐
	var recommendArticles []types.ContentSimpleView
	l.svcCtx.DB.Model(&model.Article{}).Select("id", "title", "coverUrl", "likeCount", "createTime").
		Where("userId = ?", userId).Order("likeCount+commentCount+favorCount DESC").Limit(5).Scan(&recommendArticles)
	// 视频推荐
	var recommendVideos []types.ContentSimpleView
	l.svcCtx.DB.Model(&model.Video{}).Select("id", "title", "coverUrl", "likeCount", "createTime").
		Where("userId = ?", userId).Order("likeCount+commentCount+favorCount DESC").Limit(5).Scan(&recommendVideos)

	return &types.UserHomeListResp{
		RecentArticleList:    recentArticles,
		RecentVideoList:      recentVideos,
		RecommendArticleList: recommendArticles,
		RecommendVideoList:   recommendVideos,
	}, nil
}
