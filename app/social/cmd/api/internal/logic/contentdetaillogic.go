package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContentDetailLogic {
	return &ContentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContentDetailLogic) ContentDetail(req *types.ContentDetailReq) (resp *types.ContentDetailResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	var content types.ContentView
	l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", req.Id).Scan(&content)

	// 用户信息
	var userInfoView types.UserInfoView
	userId := content.UserId
	info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId})
	_ = copier.Copy(&userInfoView, &info)
	content.UserInfo = userInfoView

	// 未登录
	if loginUserId == 0 {
		content.IsFavored = false
		content.IsLiked = false
	}

	// 是否点赞
	var isLiked bool
	l.svcCtx.DB.Model(&model.Like{}).Select("likedStatus").Where("userId = ? and itemId = ?", loginUserId, content.Id).Scan(&isLiked)
	content.IsLiked = isLiked

	// 是否收藏
	var favor model.Favor
	if err := l.svcCtx.DB.Model(&model.Favor{}).Where("userId = ? and itemType = ? and itemId = ?", loginUserId, content.Id).First(&favor).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			content.IsFavored = false
		}
	} else {
		content.IsFavored = true
	}

	return &types.ContentDetailResp{
		ContentDetail: content,
	}, nil
}
