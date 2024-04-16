package logic

import (
	"context"
	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/api/internal/types"
	"travel/app/intelligence/cmd/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type StrategyDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStrategyDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StrategyDetailLogic {
	return &StrategyDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StrategyDetailLogic) StrategyDetail(req *types.StrategyDeleteReq) (resp *types.StrategyDetailResp, err error) {
	var strategy types.StrategyView
	l.svcCtx.DB.Model(&model.Strategy{}).Where("id = ?", req.Id).Scan(&strategy)

	return &types.StrategyDetailResp{
		Strategy: strategy,
	}, nil
}
