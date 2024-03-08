package logic

import (
	"context"
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
		l.svcCtx.DB.Delete(&model.Work{}, "id = ?", req.Id)
		break
	case enum.Sell:
		l.svcCtx.DB.Model(&model.Work{}).Where("id = ?", req.Id).Update("status", enum.OnSale)
		break
	case enum.Buy:
		l.svcCtx.DB.Model(&model.Work{}).Where("id = ?", req.Id).Update("status", enum.Sold)
		// TODO:考虑使用消息队列
		// 将视频的拥有者更改为当前用户

		// 增加交易记录
		var userId int64
		l.svcCtx.DB.Model(&model.Work{}).Select("userId").Where("id = ?", req.Id).Scan(&userId)
		record := &model.Record{
			WorkId:    req.Id,
			OldUserId: userId,
			NewUserId: loginUserId,
		}
		l.svcCtx.DB.Create(record)
		break
	default:
		break
	}

	return nil
}
