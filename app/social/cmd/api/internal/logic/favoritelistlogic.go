package logic

import (
	"context"
	"travel/app/social/cmd/model"
	"travel/common/enum"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

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
	l.svcCtx.DB.Model(&model.Favorite{}).Where("userId = ?", req.UserId).Scan(&favorites)

	// 获取封面
	for i, f := range favorites {
		var favor model.Favor
		coverUrl := "默认 CoverUrl"
		l.svcCtx.DB.Model(&model.Favor{}).Where("favoriteId = ?", f.Id).Order("createTime DESC").First(&favor)
		if favor == (model.Favor{}) {
			favorites[i].CoverUrl = coverUrl
			continue
		}
		switch enum.FileType(favor.ItemType) {
		case enum.Text:
			l.svcCtx.DB.Model(&model.Article{}).Select("coverUrl").Scan(&coverUrl)
			favorites[i].CoverUrl = coverUrl
			break
		case enum.Video:
			l.svcCtx.DB.Model(&model.Video{}).Select("coverUrl").Scan(&coverUrl)
			favorites[i].CoverUrl = coverUrl
			break
		default:
			break
		}
	}

	return &types.FavoriteListResp{
		List: favorites,
	}, nil
}
