package logic

import (
	"context"
	"travel/app/social/cmd/model"
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
	userId := req.UserId
	offset := (req.PageNum - 1) * req.PageSize
	var contents []types.ContentSimpleView
	var total int64
	switch enum.FileType(req.ItemType) {
	case enum.Text:
		l.svcCtx.DB.Model(&model.Article{}).Where("userId = ?", userId).Count(&total)
		contents, _ = l.getSortedArticleList(req.SortType, offset, req.PageSize, userId)
		break
	case enum.Video:
		l.svcCtx.DB.Model(&model.Video{}).Where("userId = ?", userId).Count(&total)
		contents, _ = l.getSortedVideoList(req.SortType, offset, req.PageSize, userId)
		break
	default:
		break
	}

	return &types.UserHomeContentListResp{
		List:  contents,
		Total: int(total),
	}, nil
}

func (l *UserHomeContentListLogic) getSortedArticleList(sortType, offset, pageSize int, userId int64) ([]types.ContentSimpleView, error) {
	var contents []types.ContentSimpleView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Article{}).Offset(offset).Limit(pageSize).Order("createTime DESC").
			Where("userId = ?", userId).Scan(&contents)
		break
	case enum.Popular:
		// TODO: 考虑评论的正负性
		l.svcCtx.DB.Model(&model.Article{}).Offset(offset).Limit(pageSize).Order("likeCount+commentCount+favorCount DESC").
			Where("userId = ?", userId).Scan(&contents)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Article{}).Offset(offset).Limit(pageSize).Order("createTime ASC").
			Where("userId = ?", userId).Scan(&contents)
		break
	default:
		break
	}
	return contents, nil
}

func (l *UserHomeContentListLogic) getSortedVideoList(sortType, offset, pageSize int, userId int64) ([]types.ContentSimpleView, error) {
	var contents []types.ContentSimpleView
	switch enum.SortType(sortType) {
	case enum.Newest:
		l.svcCtx.DB.Model(&model.Video{}).Offset(offset).Limit(pageSize).Order("createTime DESC").
			Where("userId = ?", userId).Scan(&contents)
		break
	case enum.Popular:
		// TODO: 考虑评论的正负性
		l.svcCtx.DB.Model(&model.Video{}).Offset(offset).Limit(pageSize).Order("likeCount+commentCount+favorCount DESC").
			Where("userId = ?", userId).Scan(&contents)
		break
	case enum.Oldest:
		l.svcCtx.DB.Model(&model.Video{}).Offset(offset).Limit(pageSize).Order("createTime ASC").
			Where("userId = ?", userId).Scan(&contents)
		break
	default:
		break
	}
	return contents, nil
}
