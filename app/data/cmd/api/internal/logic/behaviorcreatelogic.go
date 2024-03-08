package logic

import (
	"context"
	"travel/app/data/cmd/api/internal/svc"
	"travel/app/data/cmd/api/internal/types"
	"travel/app/data/cmd/model"
	"travel/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type BehaviorCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBehaviorCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BehaviorCreateLogic {
	return &BehaviorCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BehaviorCreateLogic) BehaviorCreate(req *types.BehaviorCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	behavior := model.Behavior{
		UserId:           loginUserId,
		BehaviorItemType: req.BehaviorItemType,
		BehaviorItemId:   req.BehaviorItemId,
	}
	l.svcCtx.DB.Create(behavior)

	return nil
}
