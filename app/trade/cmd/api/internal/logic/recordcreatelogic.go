package logic

import (
	"context"
	"travel/app/trade/cmd/model"
	"travel/common/ctxdata"

	"travel/app/trade/cmd/api/internal/svc"
	"travel/app/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecordCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecordCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecordCreateLogic {
	return &RecordCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecordCreateLogic) RecordCreate(req *types.RecordCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	record := &model.Record{
		WorkId:    req.WorkId,
		OldUserId: req.UserId,
		NewUserId: loginUserId,
	}
	l.svcCtx.DB.Create(record)

	return nil
}
