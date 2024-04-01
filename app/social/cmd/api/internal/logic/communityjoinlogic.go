package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityJoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityJoinLogic {
	return &CommunityJoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityJoinLogic) CommunityJoin(req *types.CommunityJoinReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	communityJoin := &model.UserCommunity{
		UserId:      loginUserId,
		Role:        req.Role,
		CommunityId: req.CommunityId,
	}
	if err := l.svcCtx.DB.Create(communityJoin).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "加入失败")
	}

	// 增加社区成员量
	if err := l.svcCtx.DB.Model(&model.Community{}).
		Where("id = ?", req.CommunityId).
		Update("member_count", gorm.Expr("member_count + ?", 1)).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "增加失败")
	}

	return nil
}
