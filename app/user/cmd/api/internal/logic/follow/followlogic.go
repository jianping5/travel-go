package follow

import (
	"context"
	"github.com/pkg/errors"
	"travel/app/user/cmd/model"
	"travel/common/ctxdata"
	"travel/common/xerr"

	"travel/app/user/cmd/api/internal/svc"
	"travel/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var FollowDeleteError = xerr.NewErrMsg("删除失败")

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	switch followType := req.Type; followType {
	case false:
		// 关注
		follow := &model.Follow{
			UserId:       loginUserId,
			FollowUserId: req.Id,
		}
		l.svcCtx.DB.Create(follow)
	case true:
		// 取消关注
		err := l.svcCtx.DB.Delete(&model.Follow{}, "userId = ? and followUserId = ?", loginUserId, req.Id)
		if err != nil {
			return errors.Wrap(FollowDeleteError, "删除失败")
		}
	default:
		return errors.Wrap(xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR), "请求参数错误")
	}

	return nil
}
