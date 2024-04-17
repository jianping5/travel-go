package logic

import (
	"context"
	"travel/app/social/cmd/rpc/social"
	"travel/app/trade/cmd/api/internal/svc"
	"travel/app/trade/cmd/api/internal/types"
	"travel/app/trade/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserWorkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserWorkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserWorkListLogic {
	return &UserWorkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserWorkListLogic) UserWorkList(req *types.UserWorkListReq) (resp *types.UserWorkListResp, err error) {
	userId := req.UserId
	// 若 userId 为 0，则表示查看当前用户的作品列表
	if userId == 0 {
		userId = ctxdata.GetUidFromCtx(l.ctx)
	}
	offset := (req.PageNum - 1) * req.PageSize
	var works []types.WorkView
	var total int64
	tx := l.svcCtx.DB.Model(&model.Work{}).Where("user_id = ?", userId)

	// 记录总数
	countTx := tx
	countTx.Count(&total)

	tx.Offset(offset).Limit(req.PageSize).Scan(&works)

	for i, w := range works {
		info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: w.UserId, LoginUserId: userId})
		works[i].Avatar = info.Avatar
		works[i].Account = info.Account

		simple, _ := l.svcCtx.SocialRpc.ContentSimple(l.ctx, &social.ContentSimpleReq{Id: w.CopyrightId})
		works[i].Title = simple.Title
		works[i].CoverUrl = simple.CoverUrl
		works[i].ItemType = int(simple.ItemType)
	}

	return &types.UserWorkListResp{
		List:  works,
		Total: int(total),
	}, nil

	return
}
