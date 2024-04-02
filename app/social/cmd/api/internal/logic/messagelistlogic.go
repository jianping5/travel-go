package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageListLogic {
	return &MessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageListLogic) MessageList() (resp *types.MessageListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var messages []model.Message
	l.svcCtx.DB.Model(&model.Message{}).Where("user_id = ?", loginUserId).Scan(&messages)

	var messageViews []types.MessageView
	for _, m := range messages {
		var messageView types.MessageView
		_ = copier.Copy(&messageView, &m)
		// 用户信息
		userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: m.MessageUserId})
		messageView.Account = userInfo.Account
		// 内容信息
		var content model.Content
		l.svcCtx.DB.Model(&model.Content{}).Select("title", "cover_url").Where("id = ?", m.ItemId).First(&content)
		messageView.Title = content.Title
		messageView.CoverUrl = content.CoverUrl
		messageViews = append(messageViews, messageView)
	}

	return &types.MessageListResp{
		List: messageViews,
	}, nil
}
