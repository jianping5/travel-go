package logic

import (
	"context"
	"travel/app/social/cmd/model"
	"travel/common/enum"

	"travel/app/social/cmd/rpc/internal/svc"
	"travel/app/social/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentSimpleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewContentSimpleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentSimpleLogic {
	return &ContentSimpleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ContentSimpleLogic) ContentSimple(in *pb.ContentSimpleReq) (*pb.ContentSimpleResp, error) {
	var copyright model.Copyright
	l.svcCtx.DB.Model(&model.Copyright{}).Select("item_type", "item_id").Where("id = ?", in.Id).Scan(&copyright)

	var contentSimple pb.ContentSimpleResp
	var content model.Content
	l.svcCtx.DB.Model(&model.Content{}).Select("title", "cover_url", "item_type").Where("id = ?", copyright.ItemId).Scan(&content)
	contentSimple.Title = content.Title
	contentSimple.CoverUrl = content.CoverUrl
	contentSimple.ItemType = int32(enum.ARTICLE)

	return &contentSimple, nil
}
