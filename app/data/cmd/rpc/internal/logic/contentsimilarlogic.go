package logic

import (
	"context"
	"travel/app/data/cmd/model"

	"travel/app/data/cmd/rpc/internal/svc"
	"travel/app/data/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentSimilarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewContentSimilarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentSimilarLogic {
	return &ContentSimilarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ContentSimilarLogic) ContentSimilar(in *pb.ContentSimilarReq) (*pb.ContentSimilarResp, error) {
	tags := in.Tag
	if tags == nil || len(tags) == 0 {
		return &pb.ContentSimilarResp{}, nil
	}
	var itemIds []int64
	l.svcCtx.DB.Model(&model.ContentTag{}).Select("item_id").Where("item_type = ? and tag in (?)", in.ItemType, tags).Limit(10).Scan(&itemIds)

	return &pb.ContentSimilarResp{
		ItemIds: itemIds,
	}, nil
}
