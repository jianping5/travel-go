package logic

import (
	"context"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"

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
	tx := l.svcCtx.DB.Model(&model.History{}).Where("user_id = ?", loginUserId)

	// 记录总数
	countTx := tx
	countTx.Count(&total)

	// 查询历史记录
	tx.Offset(offset).Limit(req.PageSize).Scan(&historys)

	for i, h := range historys {
		// 内容信息 + 用户信息
		var content model.Content
		l.svcCtx.DB.Select("title", "cover_url", "like_count", "user_id").Where("id = ?", h.ItemId).First(&content)
		historys[i].Title = content.Title
		historys[i].CoverUrl = content.CoverUrl
		historys[i].LikeCount = content.LikeCount
		info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: content.UserId, LoginUserId: loginUserId})
		historys[i].Account = info.Account
	}

	return &types.HistoryListResp{
		List:  historys,
		Total: int(total),
	}, nil
}
