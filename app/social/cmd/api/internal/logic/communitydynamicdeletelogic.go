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

type CommunityDynamicDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityDynamicDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityDynamicDeleteLogic {
	return &CommunityDynamicDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityDynamicDeleteLogic) CommunityDynamicDelete(req *types.CommunityDynamicDeleteReq) error {
	// 判断是否有删除权限
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var dynamic model.Dynamic
	if affected := l.svcCtx.DB.Model(&model.Community{}).Select("userId").First(&dynamic).RowsAffected; affected == 0 {
		return errors.Wrap(xerr.NewErrMsg("该动态不存在"), "该动态不存在")
	}
	if dynamic.UserId != loginUserId {
		return errors.Wrap(xerr.NewErrMsg("没有权限删除"), "没有权限删除")
	}
	// 删除
	if err := l.svcCtx.DB.Delete(&model.Dynamic{}, req.Id).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	return nil
}
