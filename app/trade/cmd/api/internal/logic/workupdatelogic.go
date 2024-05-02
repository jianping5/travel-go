package logic

import (
	"context"
	"travel/app/social/cmd/rpc/social"
	"travel/app/trade/cmd/model"
	"travel/common/ctxdata"
	"travel/common/enum"

	"travel/app/trade/cmd/api/internal/svc"
	"travel/app/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWorkUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkUpdateLogic {
	return &WorkUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WorkUpdateLogic) WorkUpdate(req *types.WorkUpdateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	switch enum.WorkUpdateType(req.Type) {
	case enum.Remove:
		l.svcCtx.DB.Model(&model.Work{}).Where("id = ?", req.Id).Update("status", enum.Created)
		break
	case enum.Sell:
		l.svcCtx.DB.Model(&model.Work{}).Where("id = ?", req.Id).Update("status", enum.OnSale)
		break
	case enum.Buy:
		l.svcCtx.DB.Model(&model.Work{}).Where("id = ?", req.Id).Update("status", enum.Sold)
		// 获取版权 Id
		var copyrightId int64
		l.svcCtx.DB.Model(&model.Work{}).Select("copyright_id").Where("id = ?", req.Id).Scan(&copyrightId)
		// TODO:考虑使用消息队列
		// TODO：区块链上怎么操作呢？
		// 更改版权归属权
		l.svcCtx.SocialRpc.CopyrightUpdate(l.ctx, &social.CopyrightUpdateReq{CopyrightId: copyrightId, UserId: loginUserId, AccountAddress: req.AccountAddress})

		// 将原作品的 userId 更改为 购买用户的 userId
		l.svcCtx.SocialRpc.ContentUpdate(l.ctx, &social.ContentUpdateReq{CopyrightId: copyrightId, UserId: loginUserId})

		// 增加交易记录
		var userId int64
		l.svcCtx.DB.Model(&model.Work{}).Select("user_id").Where("id = ?", req.Id).Scan(&userId)
		record := &model.Record{
			WorkId:            req.Id,
			CopyrightId:       copyrightId,
			OldUserId:         userId,
			OldAccountAddress: req.OldAccountAddress,
			NewUserId:         loginUserId,
			NewAccountAddress: req.AccountAddress,
		}
		l.svcCtx.DB.Create(record)
		break
	default:
		break
	}

	return nil
}
