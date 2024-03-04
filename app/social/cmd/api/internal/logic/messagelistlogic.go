package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

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
	l.svcCtx.DB.Where("userId = ?", loginUserId).Scan(&messages)

	var messageViews []types.MessageView
	for _, m := range messages {
		var messageView types.MessageView
		_ = copier.Copy(&messageView, &m)
		// 用户信息
		userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: m.MessageUserId})
		messageView.Account = userInfo.Account
		// 内容信息
		switch enum.FileType(m.ItemType) {
		case enum.Text:
			var article model.Article
			l.svcCtx.DB.Model(&model.Article{}).Select("title", "coverUrl").Where("id = ?", m.ItemId).First(&article)
			messageView.Title = article.Title
			messageView.CoverUrl = article.CoverUrl
			break
		case enum.Video:
			var video model.Video
			l.svcCtx.DB.Model(&model.Video{}).Select("title", "coverUrl").Where("id = ?", m.ItemId).First(&video)
			messageView.Title = video.Title
			messageView.CoverUrl = video.CoverUrl
			break
		default:
			break
		}
		messageViews = append(messageViews, messageView)
	}

	return &types.MessageListResp{
		List: messageViews,
	}, nil
}
