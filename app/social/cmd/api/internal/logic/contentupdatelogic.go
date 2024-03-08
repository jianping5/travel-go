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

type ContentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentUpdateLogic {
	return &ContentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentUpdateLogic) ContentUpdate(req *types.ContentUpdateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var userId int64
	l.svcCtx.DB.Model(&model.Content{}).Select("userId").Where("id = ?", req.Id).Scan(&userId)
	if loginUserId != userId {
		return errors.Wrap(xerr.NewErrMsg("没有权限修改"), "没有权限修改")
	}
	// 修改内容
	if err := l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", req.Id).Updates(req).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "修改失败")
	}
	return nil
}
