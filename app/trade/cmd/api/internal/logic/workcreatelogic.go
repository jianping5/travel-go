package logic

import (
	"context"
	"github.com/pkg/errors"
	"travel/app/social/cmd/rpc/social"
	"travel/app/trade/cmd/model"
	"travel/common/ctxdata"
	"travel/common/enum"
	"travel/common/xerr"

	"travel/app/trade/cmd/api/internal/svc"
	"travel/app/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWorkCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkCreateLogic {
	return &WorkCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WorkCreateLogic) WorkCreate(req *types.WorkCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)

	// todo: 校验该版权是否属于当前用户
	check, _ := l.svcCtx.SocialRpc.CopyrightCheck(l.ctx, &social.CopyrightCheckReq{
		UserId:      loginUserId,
		CopyrightId: req.CopyrightId,
	})
	isBelonged := check.IsBelonged
	// 若不属于，则返回权限不足
	if !isBelonged {
		return errors.Wrap(xerr.NewErrMsg("权限不足"), "权限不足")
	}

	work := &model.Work{
		UserId:      loginUserId,
		CopyrightId: req.CopyrightId,
		Price:       req.Price,
		Status:      int(enum.Created),
	}
	l.svcCtx.DB.Create(work)

	return nil
}
