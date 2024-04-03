package logic

import (
	"context"
	"fmt"
	"travel/app/data/cmd/model"

	"travel/app/data/cmd/rpc/internal/svc"
	"travel/app/data/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentTagCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewContentTagCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentTagCreateLogic {
	return &ContentTagCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ContentTagCreateLogic) ContentTagCreate(in *pb.ContentTagCreateReq) (*pb.ContentTagCreateResp, error) {
	names := in.Name
	var contentTags []model.ContentTag
	for _, n := range names {
		var contentTag model.ContentTag
		contentTag.Name = n
		contentTag.ItemType = int(in.ItemType)
		contentTag.ItemId = in.ItemId
		contentTags = append(contentTags, contentTag)
	}
	// todo:test
	fmt.Println(contentTags)
	l.svcCtx.DB.Create(contentTags)

	return &pb.ContentTagCreateResp{}, nil
}
