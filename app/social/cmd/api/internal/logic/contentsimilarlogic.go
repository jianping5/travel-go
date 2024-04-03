package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/data/cmd/rpc/data"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentSimilarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentSimilarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentSimilarLogic {
	return &ContentSimilarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentSimilarLogic) ContentSimilar(req *types.ContentSimilarReq) (resp *types.ContentSimilarResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	similar, err := l.svcCtx.DataRpc.ContentSimilar(l.ctx, &data.ContentSimilarReq{
		ItemType: int32(req.ItemType),
		ItemId:   req.ItemId,
		Tag:      req.Tag,
	})
	var contents []types.ContentView
	switch enum.ItemType(req.ItemType) {
	case enum.ARTICLE:
		l.svcCtx.DB.Model(&model.Content{}).Where("id IN (?)", similar.ItemIds).Scan(&contents)
		for i, c := range contents {
			var userInfoView types.UserInfoView
			userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: c.UserId})
			_ = copier.Copy(&userInfoView, &userInfo)
			contents[i].UserInfo = userInfoView

			// 是否点赞
			var isLiked bool
			l.svcCtx.DB.Model(&model.Like{}).Select("liked_status").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.ARTICLE, c.Id).Scan(&isLiked)
			contents[i].IsLiked = isLiked

			// 是否收藏
			var favor model.Favor
			if err := l.svcCtx.DB.Model(&model.Favor{}).Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.ARTICLE, c.Id).First(&favor).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					contents[i].IsFavored = false
				}
			} else {
				contents[i].IsFavored = true
			}
		}
		break
	case enum.VIDEO:
		l.svcCtx.DB.Model(&model.Content{}).Where("id IN (?)", similar.ItemIds).Scan(&contents)
		for i, c := range contents {
			var userInfoView types.UserInfoView
			userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: c.UserId})
			_ = copier.Copy(&userInfoView, &userInfo)
			contents[i].UserInfo = userInfoView

			// 是否点赞
			var isLiked bool
			l.svcCtx.DB.Model(&model.Like{}).Select("liked_status").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.VIDEO, c.Id).Scan(&isLiked)
			contents[i].IsLiked = isLiked

			// 是否收藏
			var favor model.Favor
			if err := l.svcCtx.DB.Model(&model.Favor{}).Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.VIDEO, c.Id).First(&favor).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					contents[i].IsFavored = false
				}
			} else {
				contents[i].IsFavored = true
			}
		}
		break
	default:
		break
	}

	return &types.ContentSimilarResp{
		List: contents,
	}, nil
}
