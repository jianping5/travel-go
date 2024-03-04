package logic

import (
	"context"
	"github.com/pkg/errors"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryDeleteLogic {
	return &HistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HistoryDeleteLogic) HistoryDelete(req *types.HistoryDeleteReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	id := req.Id
	// 清空历史记录
	if id == 0 {
		if err := l.svcCtx.DB.Delete(&model.History{}, "userId = ?", loginUserId).Error; err != nil {
			return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
		}
	}
	// 删除指定历史记录
	if err := l.svcCtx.DB.Delete(&model.History{}, "id = ?", id).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	return nil
}
