package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
	"travel/common/enum"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityCreateLogic {
	return &CommunityCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityCreateLogic) CommunityCreate(req *types.CommunityCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	community := &model.Community{
		UserId:      loginUserId,
		Name:        req.Name,
		Description: req.Description,
		Avatar:      "https://uploadfiles.nowcoder.com/images/20211202/39500431_1638437021506/9C4ED8A2DD30EE49527DB5C02BC4F68B",
	}
	if err := l.svcCtx.DB.Create(community).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "社区创建失败")
	}

	// 增加用户社区关系
	communityJoin := &model.UserCommunity{
		UserId:      loginUserId,
		Role:        int(enum.Operator),
		CommunityId: community.Id,
	}
	if err := l.svcCtx.DB.Create(communityJoin).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "加入失败")
	}

	// 增加社区成员量
	if err := l.svcCtx.DB.Model(&model.Community{}).
		Where("id = ?", community.Id).
		Update("member_count", gorm.Expr("member_count + ?", 1)).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "增加失败")
	}

	return nil
}
