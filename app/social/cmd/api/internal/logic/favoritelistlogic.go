package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
	"travel/common/enum"

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
	userId := req.UserId
	// 这里的 itemId 主要是在收藏某个具体的 item 时会需要，其它的传 0 即可
	// todo：怎么更改成不传也可也不报错？
	itemId := req.ItemId
	// 传输 0，则表示获取自己的
	if userId == 0 {
		userId = ctxdata.GetUidFromCtx(l.ctx)
	}

	var favorites []types.FavoriteListView
	l.svcCtx.DB.Model(&model.Favorite{}).Where("user_id = ?", userId).Scan(&favorites)

	// 获取封面
	for i, f := range favorites {
		var favor model.Favor
		coverUrl := "https://cdn.pixabay.com/photo/2023/12/14/00/20/alaska-8448009_1280.jpg"
		l.svcCtx.DB.Model(&model.Favor{}).Where("favorite_id = ?", f.Id).Order("create_time DESC").First(&favor)
		if favor == (model.Favor{}) {
			favorites[i].CoverUrl = coverUrl
		} else {
			l.svcCtx.DB.Model(&model.Content{}).Select("cover_url").Where("id = ?", favor.ItemId).Scan(&coverUrl)
			favorites[i].CoverUrl = coverUrl
		}
		// 是否收藏
		var favorModel model.Favor
		if err := l.svcCtx.DB.Model(&model.Favor{}).Where("user_id = ? and item_type = ? and item_id = ?", userId, enum.VIDEO, itemId).First(&favorModel).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				favorites[i].IsFavored = false
			}
		} else {
			favorites[i].IsFavored = true
		}
	}

	return &types.FavoriteListResp{
		List: favorites,
	}, nil
}
