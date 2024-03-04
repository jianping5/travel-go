package logic

import (
	"context"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryListLogic {
	return &HistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HistoryListLogic) HistoryList(req *types.HistoryListReq) (resp *types.HistoryListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	offset := (req.PageNum - 1) * req.PageSize
	var total int64
	var historys []types.HistoryView
	tx := l.svcCtx.DB.Model(&model.History{}).Where("userId = ?", loginUserId)

	// 记录总数
	countTx := tx
	countTx.Count(&total)

	// 查询历史记录
	tx.Offset(offset).Limit(req.PageSize).Scan(&historys)

	for i, h := range historys {
		// 内容信息 + 用户信息
		switch enum.FileType(h.ItemType) {
		case enum.Text:
			var article model.Article
			l.svcCtx.DB.Select("title", "coverUrl", "likeCount", "userId").Where("id = ?", h.ItemId).First(&article)
			historys[i].Title = article.Title
			historys[i].CoverUrl = article.CoverUrl
			historys[i].LikeCount = article.LikeCount
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: article.UserId})
			historys[i].Account = info.Account
			break
		case enum.Video:
			var video model.Video
			l.svcCtx.DB.Select("title", "coverUrl", "likeCount").Where("id = ?", h.ItemId).First(&video)
			historys[i].Title = video.Title
			historys[i].CoverUrl = video.CoverUrl
			historys[i].LikeCount = video.LikeCount
			info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: video.UserId})
			historys[i].Account = info.Account
			break
		default:
			break
		}
	}

	return &types.HistoryListResp{
		List:  historys,
		Total: int(total),
	}, nil
}
