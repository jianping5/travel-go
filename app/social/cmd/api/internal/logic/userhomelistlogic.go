package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	userId := req.UserId
	// 最新文章
	var recentArticles []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "user_id", "title", "cover_url", "like_count", "create_time").
		Where("user_id = ? and item_type = ?", userId, enum.ARTICLE).Order("create_time DESC").Limit(5).Scan(&recentArticles)
	l.SetContentInfo(&recentArticles, loginUserId, enum.ARTICLE)
	// 最新视频
	var recentVideos []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "user_id", "title", "cover_url", "like_count", "create_time").
		Where("user_id = ? and item_type = ?", userId, enum.VIDEO).Order("create_time DESC").Limit(5).Scan(&recentVideos)
	l.SetContentInfo(&recentVideos, loginUserId, enum.VIDEO)
	// TODO: 推荐是否考虑评论的正负性
	// 文章推荐
	var recommendArticles []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "user_id", "title", "cover_url", "like_count", "create_time").
		Where("user_id = ? and item_type = ?", userId, enum.ARTICLE).Order("like_count+comment_count+favor_count DESC").Limit(5).Scan(&recommendArticles)
	l.SetContentInfo(&recommendArticles, loginUserId, enum.ARTICLE)
	// 视频推荐
	var recommendVideos []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "user_id", "title", "cover_url", "like_count", "create_time").
		Where("user_id = ? and item_type = ?", userId, enum.VIDEO).Order("like_count+comment_count+favor_count DESC").Limit(5).Scan(&recommendVideos)
	l.SetContentInfo(&recommendVideos, loginUserId, enum.VIDEO)

	return &types.UserHomeListResp{
		RecentArticleList:    recentArticles,
		RecentVideoList:      recentVideos,
		RecommendArticleList: recommendArticles,
		RecommendVideoList:   recommendVideos,
	}, nil
}

func (l *UserHomeListLogic) SetContentInfo(contents *[]types.ContentView, loginUserId int64, itemType enum.ItemType) {
	for i, v := range *contents {
		// 用户信息
		var userInfoView types.UserInfoView
		userId := v.UserId
		info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId})
		_ = copier.Copy(&userInfoView, &info)
		(*contents)[i].UserInfo = userInfoView

		// 是否点赞
		var isLiked bool
		l.svcCtx.DB.Model(&model.Like{}).Select("liked_status").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, itemType, v.Id).Scan(&isLiked)
		(*contents)[i].IsLiked = isLiked

		// 是否收藏
		var favor model.Favor
		if err := l.svcCtx.DB.Model(&model.Favor{}).Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, itemType, v.Id).First(&favor).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				(*contents)[i].IsFavored = false
			}
		} else {
			(*contents)[i].IsFavored = true
		}
	}
}
