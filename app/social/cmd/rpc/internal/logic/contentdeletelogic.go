package logic

import (
	"context"
	"travel/app/social/cmd/model"

	"travel/app/social/cmd/rpc/internal/svc"
	"travel/app/social/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewContentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentDeleteLogic {
	return &ContentDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ContentDeleteLogic) ContentDelete(in *pb.ContentDeleteReq) (*pb.ContentDeleteResp, error) {
	// todo: add your logic here and delete this line
	var contentId int64
	l.svcCtx.DB.Model(&model.Copyright{}).Select("item_id").Where("id = ?", in.Id).Scan(&contentId)
	l.svcCtx.DB.Delete(&model.Content{}, "id = ?", contentId)

	return &pb.ContentDeleteResp{}, nil
}
