package logic

import (
	"context"
	"travel/app/social/cmd/model"

	"travel/app/social/cmd/rpc/internal/svc"
	"travel/app/social/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewContentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentUpdateLogic {
	return &ContentUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ContentUpdateLogic) ContentUpdate(in *pb.ContentUpdateReq) (*pb.ContentUpdateResp, error) {
	// todo: add your logic here and delete this line
	var contentId int64
	l.svcCtx.DB.Model(&model.Copyright{}).Select("item_id").Where("id = ?", in.CopyrightId).Scan(&contentId)
	l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", contentId).Update("user_id = ?", in.UserId)

	return &pb.ContentUpdateResp{}, nil
}
