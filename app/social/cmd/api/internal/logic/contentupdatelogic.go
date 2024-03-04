package logic

import (
	"context"
	"github.com/pkg/errors"
	"travel/app/social/cmd/model"
	"travel/common/ctxdata"
	"travel/common/enum"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentUpdateLogic {
	return &ContentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentUpdateLogic) ContentUpdate(req *types.ContentUpdateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	itemType := req.ItemType
	switch enum.FileType(itemType) {
	case enum.Text:
		var userId int64
		l.svcCtx.DB.Model(&model.Article{}).Select("userId").Where("id = ?", req.ItemId).Scan(&userId)
		if loginUserId != userId {
			return errors.Wrap(xerr.NewErrMsg("没有权限修改"), "没有权限修改")
		}
		// 修改文章
		if err := l.svcCtx.DB.Model(&model.Article{}).Where("id = ?", req.ItemId).Updates(req).Error; err != nil {
			return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "修改失败")
		}
		break
	case enum.Video:
		var userId int64
		l.svcCtx.DB.Model(&model.Video{}).Select("userId").Where("id = ?", req.ItemId).Scan(&userId)
		if loginUserId != userId {
			return errors.Wrap(xerr.NewErrMsg("没有权限修改"), "没有权限修改")
		}
		// 修改视频
		if err := l.svcCtx.DB.Model(&model.Video{}).Where("id = ?", req.ItemId).Updates(req).Error; err != nil {
			return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "修改失败")
		}
		break
	default:
		return errors.Wrap(xerr.NewErrMsg("参数错误"), "参数错误")
	}
	return nil
}
