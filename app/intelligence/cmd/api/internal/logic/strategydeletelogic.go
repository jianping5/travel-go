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

type StrategyDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStrategyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StrategyDeleteLogic {
	return &StrategyDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StrategyDeleteLogic) StrategyDelete(req *types.StrategyDeleteReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var userId int64
	l.svcCtx.DB.Model(&model.Strategy{}).Select("userId").Where("id = ?", req.Id).Scan(&userId)
	if loginUserId != userId {
		return errors.Wrap(xerr.NewErrMsg("权限不足"), "权限不足")
	}

	l.svcCtx.DB.Delete(&model.Strategy{}, "id = ?", req.Id)

	return nil
}
