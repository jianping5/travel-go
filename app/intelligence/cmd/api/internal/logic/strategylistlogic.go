package logic

import (
	"context"
	"travel/app/intelligence/cmd/model"
	"travel/common/ctxdata"

	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StrategyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStrategyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StrategyListLogic {
	return &StrategyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StrategyListLogic) StrategyList() (resp *types.StrategyListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)

	var strategyList []types.StrategyView
	l.svcCtx.DB.Model(&model.Strategy{}).Where("userId = ?", loginUserId).Scan(&strategyList)

	return &types.StrategyListResp{
		List: strategyList,
	}, nil
}
