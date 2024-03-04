package logic

import (
	"context"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/enum"
	"travel/common/tool"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavorListLogic {
	return &FavorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavorListLogic) FavorList(req *types.FavorListReq) (resp *types.FavorListResp, err error) {
	var favors []types.FavorView
	l.svcCtx.DB.Model(&model.Favor{}).Where("favoriteId = ?", req.FavoriteId).Scan(&favors)
	itemType := req.ItemType
	// 注入内容信息+用户信息
	for i, f := range favors {
		switch enum.FileType(itemType) {
		case enum.Text:
			var article model.Article
			l.svcCtx.DB.Select("title", "coverUrl", "likeCount", "userId", "createTime").Where("id = ?", f.ItemId).First(&article)
			favors[i].Title = article.Title
			favors[i].CoverUrl = article.CoverUrl
			favors[i].LikeCount = article.LikeCount
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: article.UserId})
			favors[i].Account = info.Account
			favors[i].CreateTime = tool.TimeToString(article.CreateTime)
			break
		case enum.Video:
			var video model.Video
			l.svcCtx.DB.Select("title", "coverUrl", "likeCount", "userId", "createTime").Where("id = ?", f.ItemId).First(&video)
			favors[i].Title = video.Title
			favors[i].CoverUrl = video.CoverUrl
			favors[i].LikeCount = video.LikeCount
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: video.UserId})
			favors[i].Account = info.Account
			favors[i].CreateTime = tool.TimeToString(video.CreateTime)
			break
		default:
			break
		}
	}

	return &types.FavorListResp{
		List: favors,
	}, nil
}
