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

type CommunityDynamicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityDynamicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityDynamicCreateLogic {
	return &CommunityDynamicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityDynamicCreateLogic) CommunityDynamicCreate(req *types.CommunityDynamicCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	dynamic := &model.Dynamic{
		UserId:       loginUserId,
		Title:        req.Title,
		CommunityId:  req.CommunityId,
		FileType:     req.FileType,
		Content:      req.Content,
		LikeCount:    0,
		CommentCount: 0,
	}
	if err := l.svcCtx.DB.Create(dynamic).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "创建失败")
	}

	return nil
}
