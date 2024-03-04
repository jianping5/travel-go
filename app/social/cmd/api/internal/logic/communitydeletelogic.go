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

type CommunityDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityDeleteLogic {
	return &CommunityDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityDeleteLogic) CommunityDelete(req *types.CommunityDeleteReq) error {
	// 判断是否有删除权限
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var community model.Community
	if affected := l.svcCtx.DB.Model(&model.Community{}).Select("userId").First(&community).RowsAffected; affected == 0 {
		return errors.Wrap(xerr.NewErrMsg("该社区不存在"), "该社区不存在")
	}
	if community.UserId != loginUserId {
		return errors.Wrap(xerr.NewErrMsg("没有权限删除"), "没有权限删除")
	}
	// 删除
	if err := l.svcCtx.DB.Delete(&model.Community{}, req.Id).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	return nil
}
