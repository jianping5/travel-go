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

type CommunityQuitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityQuitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityQuitLogic {
	return &CommunityQuitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityQuitLogic) CommunityQuit(req *types.CommunityQuitReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	if err := l.svcCtx.DB.Delete(&model.UserCommunity{}, "user_id = ? and community_id = ?", loginUserId, req.CommunityId).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "删除失败")
	}

	// 减少对应社区的成员量
	if err := l.svcCtx.DB.Model(&model.Community{}).Where("id = ?", req.CommunityId).
		Update("member_count", gorm.Expr("member_count - ?", 1)).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "减少失败")
	}

	return nil
}
