package logic

import (
	"context"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListReq) (resp *types.FavoriteListResp, err error) {
	var favorites []types.FavoriteListView
	l.svcCtx.DB.Model(&model.Favorite{}).Where("user_id = ?", req.UserId).Scan(&favorites)

	// 获取封面
	for i, f := range favorites {
		var favor model.Favor
		coverUrl := "https://cdn.pixabay.com/photo/2023/12/14/00/20/alaska-8448009_1280.jpg"
		l.svcCtx.DB.Model(&model.Favor{}).Where("favorite_id = ?", f.Id).Order("create_time DESC").First(&favor)
		if favor == (model.Favor{}) {
			favorites[i].CoverUrl = coverUrl
			continue
		}
		l.svcCtx.DB.Model(&model.Content{}).Select("cover_url").Where("id = ?", favor.ItemId).Scan(&coverUrl)
		favorites[i].CoverUrl = coverUrl
	}

	return &types.FavoriteListResp{
		List: favorites,
	}, nil
}
