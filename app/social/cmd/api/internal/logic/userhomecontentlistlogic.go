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

type UserHomeContentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomeContentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomeContentListLogic {
	return &UserHomeContentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomeContentListLogic) UserHomeContentList(req *types.UserHomeContentListReq) (resp *types.UserHomeContentListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	userId := req.UserId
	offset := (req.PageNum - 1) * req.PageSize
	var contents []types.ContentView
	var total int64
	switch enum.ItemType(req.ItemType) {
	case enum.ARTICLE:
		l.svcCtx.DB.Model(&model.Content{}).Where("userId = ? and itemType = ?", userId, enum.ARTICLE).Count(&total)
		contents, _ = l.getSortedArticleList(req.SortType, offset, req.PageSize, userId)
		for i, a := range contents {
			// 用户信息
			var userInfoView types.UserInfoView
			userId := a.UserId
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId})
			_ = copier.Copy(&userInfoView, &info)
			contents[i].UserInfo = userInfoView

			// 是否点赞
			var isLiked bool
			l.svcCtx.DB.Model(&model.Like{}).Select("likedStatus").Where("userId = ? and itemType = ? and itemId = ?", loginUserId, enum.ARTICLE, a.Id).Scan(&isLiked)
			contents[i].IsLiked = isLiked

			// 是否收藏
			var favor model.Favor
			if err := l.svcCtx.DB.Model(&model.Favor{}).Where("userId = ? and itemType = ? and itemId = ?", loginUserId, enum.ARTICLE, a.Id).First(&favor).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					contents[i].IsFavored = false
				}
			} else {
				contents[i].IsFavored = true
			}
		}
		break
	case enum.VIDEO:
		l.svcCtx.DB.Model(&model.Content{}).Where("userId = ? and itemType = ?", userId, enum.VIDEO).Count(&total)
		contents, _ = l.getSortedVideoList(req.SortType, offset, req.PageSize, userId)
		for i, v := range contents {
			// 用户信息
			var userInfoView types.UserInfoView
			userId := v.UserId
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId})
			_ = copier.Copy(&userInfoView, &info)
			contents[i].UserInfo = userInfoView

			// 是否点赞
			var isLiked bool
			l.svcCtx.DB.Model(&model.Like{}).Select("likedStatus").Where("userId = ? and itemType = ? and itemId = ?", loginUserId, enum.VIDEO, v.Id).Scan(&isLiked)
			contents[i].IsLiked = isLiked

			// 是否收藏
			var favor model.Favor
			if err := l.svcCtx.DB.Model(&model.Favor{}).Where("userId = ? and itemType = ? and itemId = ?", loginUserId, enum.VIDEO, v.Id).First(&favor).Error; err != nil {
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

	return &types.UserHomeContentListResp{
		List:  contents,
		Total: int(total),
	}, nil
}

func (l *UserHomeContentListLogic) getSortedArticleList(sortType, offset, pageSize int, userId int64) ([]types.ContentView, error) {
	var contents []types.ContentView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("createTime DESC").
			Where("userId = ? and itemType = ?", userId, enum.ARTICLE).Scan(&contents)
		break
	case enum.Popular:
		// TODO: 考虑评论的正负性
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("likeCount+commentCount+favorCount DESC").
			Where("userId = ? and itemType = ?", userId, enum.ARTICLE).Scan(&contents)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("createTime ASC").
			Where("userId = ? and itemType = ?", userId, enum.ARTICLE).Scan(&contents)
		break
	default:
		break
	}
	return contents, nil
}

func (l *UserHomeContentListLogic) getSortedVideoList(sortType, offset, pageSize int, userId int64) ([]types.ContentView, error) {
	var contents []types.ContentView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("createTime DESC").
			Where("userId = ? and itemType = ?", userId, enum.VIDEO).Scan(&contents)
		break
	case enum.Popular:
		// TODO: 考虑评论的正负性
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("likeCount+commentCount+favorCount DESC").
			Where("userId = ? and itemType = ?", userId, enum.VIDEO).Scan(&contents)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Content{}).Offset(offset).Limit(pageSize).Order("createTime ASC").
			Where("userId = ? and itemType = ?", userId, enum.VIDEO).Scan(&contents)
		break
	default:
		break
	}
	return contents, nil
}
