package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"travel/app/data/cmd/rpc/data"
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
	// 转化成 json
	tagJson, _ := json.Marshal(req.Tag)
	switch enum.ItemType(itemType) {
	case enum.ARTICLE:
		// todo：文章
		content := &model.Content{
			UserId:   loginUserId,
			ItemType: int(enum.ARTICLE),
			Title:    req.Title,
			CoverUrl: req.CoverUrl,
			Content:  req.Content,
			Tag:      tagJson,
		}
		if err := l.svcCtx.DB.Create(&content).Error; err != nil {
			return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "创建失败")
		}

		// 插入内容标签表
		l.svcCtx.DataRpc.ContentTagCreate(l.ctx, &data.ContentTagCreateReq{
			Name:     req.Tag,
			ItemType: int32(enum.VIDEO),
			ItemId:   content.Id,
		})

		// todo：给关注该用户的人发送消息
		break
	case enum.VIDEO:
		content := &model.Content{
			UserId:      loginUserId,
			ItemType:    int(enum.VIDEO),
			Title:       req.Title,
			CoverUrl:    req.CoverUrl,
			Content:     req.Content,
			Description: req.Description,
			Tag:         tagJson,
		}
		if err := l.svcCtx.DB.Create(&content).Error; err != nil {
			return errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), "创建失败")
		}

		// 插入内容标签表
		l.svcCtx.DataRpc.ContentTagCreate(l.ctx, &data.ContentTagCreateReq{
			Name:     req.Tag,
			ItemType: int32(enum.VIDEO),
			ItemId:   content.Id,
		})

		// todo：给关注该用户的人发送消息
		break
	default:
		return errors.Wrap(xerr.NewErrMsg("参数错误"), "参数错误")
	}

	return nil
}
