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

type ContentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentDeleteLogic {
	return &ContentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentDeleteLogic) ContentDelete(req *types.ContentDeleteReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	// 文章
	var userId int64
	l.svcCtx.DB.Model(&model.Content{}).Select("user_id").Where("id = ?", req.Id).Scan(&userId)
	if loginUserId != userId {
		return errors.Wrap(xerr.NewErrMsg("没有权限删除"), "没有权限删除")
	}
	// 删除
	if err := l.svcCtx.DB.Delete(&model.Content{}, "id = ?", req.Id).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	return nil
}
