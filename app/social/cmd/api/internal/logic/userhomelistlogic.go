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
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "title", "coverUrl", "likeCount", "createTime").
		Where("userId = ? and itemType = ?", userId, enum.ARTICLE).Order("createTime DESC").Limit(5).Scan(&recentArticles)
	l.SetContentInfo(&recentArticles, loginUserId, enum.ARTICLE)
	// 最新视频
	var recentVideos []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "title", "coverUrl", "likeCount", "createTime").
		Where("userId = ? and itemType = ?", userId, enum.VIDEO).Order("createTime DESC").Limit(5).Scan(&recentVideos)
	l.SetContentInfo(&recentArticles, loginUserId, enum.VIDEO)
	// TODO: 推荐是否考虑评论的正负性
	// 文章推荐
	var recommendArticles []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "title", "coverUrl", "likeCount", "createTime").
		Where("userId = ? and itemType = ?", userId, enum.ARTICLE).Order("likeCount+commentCount+favorCount DESC").Limit(5).Scan(&recommendArticles)
	l.SetContentInfo(&recentArticles, loginUserId, enum.ARTICLE)
	// 视频推荐
	var recommendVideos []types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Select("id", "title", "coverUrl", "likeCount", "createTime").
		Where("userId = ? and itemType = ?", userId, enum.VIDEO).Order("likeCount+commentCount+favorCount DESC").Limit(5).Scan(&recommendVideos)
	l.SetContentInfo(&recentArticles, loginUserId, enum.VIDEO)

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
		l.svcCtx.DB.Model(&model.Like{}).Select("likedStatus").Where("userId = ? and itemType = ? and itemId = ?", loginUserId, itemType, v.Id).Scan(&isLiked)
		(*contents)[i].IsLiked = isLiked

		// 是否收藏
		var favor model.Favor
		if err := l.svcCtx.DB.Model(&model.Favor{}).Where("userId = ? and itemType = ? and itemId = ?", loginUserId, itemType, v.Id).First(&favor).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				(*contents)[i].IsFavored = false
			}
		} else {
			(*contents)[i].IsFavored = true
		}
	}
}
