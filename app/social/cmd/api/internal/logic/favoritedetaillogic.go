package logic

import (
	"context"
	"fmt"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteDetailLogic {
	return &FavoriteDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteDetailLogic) FavoriteDetail(req *types.FavoriteDeleteReq) (resp *types.FavoriteDetailResp, err error) {
	var favorite types.FavoriteListView
	l.svcCtx.DB.Model(&model.Favorite{}).Where("id = ?", req.Id).Scan(&favorite)

	// 获取封面
	var favor model.Favor
	coverUrl := "https://cdn.pixabay.com/photo/2023/12/14/00/20/alaska-8448009_1280.jpg"
	l.svcCtx.DB.Model(&model.Favor{}).Where("favorite_id = ?", req.Id).Order("create_time DESC").First(&favor)

	// todo: test
	fmt.Println("favor", favor)
	fmt.Println(favor == (model.Favor{}))

	if favor == (model.Favor{}) {
		favorite.CoverUrl = coverUrl
	} else {
		l.svcCtx.DB.Model(&model.Content{}).Select("cover_url").Where("id = ?", favor.ItemId).Scan(&coverUrl)
		favorite.CoverUrl = coverUrl
	}

	return &types.FavoriteDetailResp{
		FavoriteDetail: favorite,
	}, nil
}
