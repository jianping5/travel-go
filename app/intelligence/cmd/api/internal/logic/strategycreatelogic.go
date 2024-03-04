package logic

import (
	"context"
	"github.com/pkg/errors"
	"travel/app/intelligence/cmd/model"
	"travel/common/ctxdata"
	"travel/common/xerr"

	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StrategyCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStrategyCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StrategyCreateLogic {
	return &StrategyCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StrategyCreateLogic) StrategyCreate(req *types.StrategyCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	strategy := &model.Strategy{
		UserId:      loginUserId,
		Destination: req.Destination,
		Duration:    req.Duration,
		Budget:      req.Budget,
		TripGroup:   req.TripGroup,
		TripMood:    req.TripMood,
		Strategy:    req.Strategy,
	}

	if err := l.svcCtx.DB.Create(strategy).Error; err != nil {
		return errors.Wrap(xerr.NewErrMsg("创建失败"), "创建失败")
	}

	return nil
}
