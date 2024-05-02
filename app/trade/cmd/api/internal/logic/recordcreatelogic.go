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
	// 获取获取对应版权 id
	var copyrightId int64
	l.svcCtx.DB.Model(&model.Work{}).Select("copyright_id").Where("id = ?", req.WorkId).Scan(&copyrightId)
	record := &model.Record{
		WorkId:      req.WorkId,
		CopyrightId: copyrightId,
		OldUserId:   req.UserId,
		NewUserId:   loginUserId,
	}
	l.svcCtx.DB.Create(record)

	return nil
}
