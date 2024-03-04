package logic

import (
	"context"

	"travel/app/intelligence/cmd/api/internal/svc"
	"travel/app/intelligence/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StrategyGenerateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStrategyGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StrategyGenerateLogic {
	return &StrategyGenerateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StrategyGenerateLogic) StrategyGenerate(req *types.StrategyGenerateReq) (resp *types.StrategyGenerateResp, err error) {
	// 获取请求参数

	// TODO：调用 AI 接口

	// 处理响应值，并返回

	return
}
