package logic

import (
	"context"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyrightCreateReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyrightCreateReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyrightCreateReqLogic {
	return &CopyrightCreateReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyrightCreateReqLogic) CopyrightCreateReq(req *types.CopyrightCreateReq) error {
	//loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	// 用户信息
	//userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: loginUserId, LoginUserId: loginUserId})
	// 作品信息
	var content model.Content
	l.svcCtx.DB.Model(&model.Content{}).Select("title", "description", "coverUrl", "createTime").Where("id = ?", req.ItemId).Scan(&content)

	// TODO: 消息队列，上传文件获取链接，版权存证到区块链，最后存入相关信息到数据库

	return nil
}
