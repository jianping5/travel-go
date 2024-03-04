package logic

import (
	"context"
	"travel/app/social/cmd/model"
	"travel/common/enum"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

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
	//userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: loginUserId})
	// 作品信息
	switch enum.FileType(req.ItemType) {
	case enum.Text:
		var article model.Article
		l.svcCtx.DB.Model(&model.Article{}).Select("title", "description", "coverUrl", "createTime").Where("id = ?", req.ItemId).Scan(&article)
		break
	case enum.Video:
		var video model.Video
		l.svcCtx.DB.Model(&model.Video{}).Select("title", "description", "coverUrl", "createTime").Where("id = ?", req.ItemId).Scan(&video)
		break
	default:
		break
	}

	// TODO: 消息队列，上传文件获取链接，版权存证到区块链，最后存入相关信息到数据库

	return nil
}
