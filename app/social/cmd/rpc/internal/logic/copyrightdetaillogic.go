package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"travel/app/social/cmd/model"
	"travel/app/social/cmd/rpc/internal/svc"
	"travel/app/social/cmd/rpc/pb/pb"
)

type CopyrightDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCopyrightDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyrightDetailLogic {
	return &CopyrightDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CopyrightDetailLogic) CopyrightDetail(in *pb.CopyrightDetailReq) (*pb.CopyrightDetailResp, error) {
	// 版权信息
	var copyright pb.CopyrightDetailResp
	l.svcCtx.DB.Model(&model.Copyright{}).Where("id = ?", in.Id).Scan(&copyright)

	var content model.Content
	l.svcCtx.DB.Model(&model.Content{}).Select("title", "cover_url", "content", "description").Where("id = ?", copyright.ItemId).Scan(&content)
	copyright.Title = content.Title
	copyright.CoverUrl = content.CoverUrl
	copyright.Content = content.Content
	copyright.Description = content.Description

	return &copyright, nil
}
