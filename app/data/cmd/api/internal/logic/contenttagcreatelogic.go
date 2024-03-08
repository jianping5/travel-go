package logic

import (
	"context"
	"travel/app/data/cmd/model"

	"travel/app/data/cmd/api/internal/svc"
	"travel/app/data/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentTagCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentTagCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentTagCreateLogic {
	return &ContentTagCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentTagCreateLogic) ContentTagCreate(req *types.ContentTagCreateReq) error {
	names := req.Name
	var contentTags []model.ContentTag
	for _, n := range names {
		var contentTag model.ContentTag
		contentTag.Name = n
		contentTag.ItemType = req.ItemType
		contentTag.ItemId = req.ItemId
		contentTags = append(contentTags, contentTag)
	}
	l.svcCtx.DB.Create(contentTags)

	return nil
}
