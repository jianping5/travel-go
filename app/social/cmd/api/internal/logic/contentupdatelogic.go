package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
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
	l.svcCtx.DB.Model(&model.Content{}).Select("user_id").Where("id = ?", req.Id).Scan(&userId)
	if loginUserId != userId {
		return errors.Wrap(xerr.NewErrMsg("没有权限修改"), "没有权限修改")
	}
	// 修改内容
	var update model.ContentUpdateReq
	_ = copier.Copy(&update, req)
	update.Tag, _ = json.Marshal(req.Tag)
	if err := l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", update.Id).Updates(&update).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "修改失败")
	}
	return nil
}
