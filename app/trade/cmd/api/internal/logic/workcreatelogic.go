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

type WorkCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWorkCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkCreateLogic {
	return &WorkCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WorkCreateLogic) WorkCreate(req *types.WorkCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	work := &model.Work{
		UserId:      loginUserId,
		CopyrightId: req.CopyrightId,
		Price:       req.Price,
		Status:      int(enum.Created),
	}
	l.svcCtx.DB.Create(work)

	return nil
}
