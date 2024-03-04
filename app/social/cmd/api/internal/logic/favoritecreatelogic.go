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

type FavoriteCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteCreateLogic {
	return &FavoriteCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteCreateLogic) FavoriteCreate(req *types.FavoriteCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	favorite := &model.Favorite{
		UserId: loginUserId,
		Name:   req.Name,
	}
	if err := l.svcCtx.DB.Create(favorite).Error; err != nil {
		return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "创建收藏夹失败")
	}

	return nil
}
