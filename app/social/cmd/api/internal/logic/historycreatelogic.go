package logic

import (
	"context"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryCreateLogic {
	return &HistoryCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HistoryCreateLogic) HistoryCreate(req *types.HistoryCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)

	history := &model.History{
		UserId: loginUserId,
		ItemId: req.ItemId,
	}
	l.svcCtx.DB.Create(history)

	return nil
}
