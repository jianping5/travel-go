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

type ContentCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentCreateLogic {
	return &ContentCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentCreateLogic) ContentCreate(req *types.ContentCreateReq) error {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	itemType := req.ItemType
	switch enum.FileType(itemType) {
	case enum.Text:
		// todo：文章
		content := &model.Article{
			UserId:   loginUserId,
			Title:    req.Title,
			CoverUrl: req.CoverUrl,
			Content:  req.Content,
			Tag:      req.Tag,
		}
		if err := l.svcCtx.DB.Create(&content).Error; err != nil {
			return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "创建失败")
		}

		// todo：给关注该用户的人发送消息
		break
	case enum.Video:
		content := &model.Video{
			UserId:      loginUserId,
			Title:       req.Title,
			CoverUrl:    req.CoverUrl,
			Content:     req.Content,
			Description: req.Description,
			Tag:         req.Tag,
		}
		if err := l.svcCtx.DB.Create(&content).Error; err != nil {
			return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "创建失败")
		}

		// todo：给关注该用户的人发送消息
		break
	default:
		return errors.Wrap(xerr.NewErrMsg("参数错误"), "参数错误")
	}

	return nil
}
