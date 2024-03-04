package logic

import (
	"context"
	"github.com/pkg/errors"
	"travel/app/social/cmd/model"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityUpdateLogic {
	return &CommunityUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityUpdateLogic) CommunityUpdate(req *types.CommunityUpdateReq) error {
	description := req.Description
	avatar := req.Avatar
	// 修改简介
	if description != "" {
		if err := l.svcCtx.DB.Model(&model.Community{}).Where("id = ?", req.Id).
			Update("description", description).Error; err != nil {
				return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "修改失败")
		}
	}
	// 修改头像
	if avatar != "" {
		if err := l.svcCtx.DB.Model(&model.Community{}).Where("id = ?", req.Id).
			Update("avatar", avatar).Error; err != nil {
			return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "修改失败")
		}
	}

	return nil
}
