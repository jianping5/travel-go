package logic

import (
	"context"
	"travel/app/social/cmd/rpc/social"
	"travel/app/trade/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"

	"travel/app/trade/cmd/api/internal/svc"
	"travel/app/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWorkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkListLogic {
	return &WorkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WorkListLogic) WorkList(req *types.WorkListReq) (resp *types.WorkListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	offset := (req.PageNum - 1) * req.PageSize
	var works []types.WorkView
	var total int64
	tx := l.svcCtx.DB.Model(&model.Work{}).Where("status = ?", enum.OnSale)

	// 记录总数
	countTx := tx
	countTx.Count(&total)

	tx.Offset(offset).Limit(req.PageSize).Scan(&works)

	for i, w := range works {
		info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: w.UserId, LoginUserId: loginUserId})
		works[i].Avatar = info.Avatar
		works[i].Account = info.Account

		simple, _ := l.svcCtx.SocialRpc.ContentSimple(l.ctx, &social.ContentSimpleReq{Id: w.CopyrightId})
		works[i].Title = simple.Title
		works[i].CoverUrl = simple.CoverUrl
		works[i].ItemType = int(simple.ItemType)
	}

	return &types.WorkListResp{
		List:  works,
		Total: int(total),
	}, nil
}
