package logic

import (
	"context"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
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
	// 注入内容信息+用户信息
	for i, f := range favors {
		var content model.Content
		l.svcCtx.DB.Select("title", "coverUrl", "likeCount", "userId", "createTime").Where("id = ?", f.ItemId).First(&content)
		favors[i].Title = content.Title
		favors[i].CoverUrl = content.CoverUrl
		favors[i].LikeCount = content.LikeCount
		info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: content.UserId})
		favors[i].Account = info.Account
		favors[i].CreateTime = tool.TimeToString(content.CreateTime)
	}

	return &types.FavorListResp{
		List: favors,
	}, nil
}
