package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"
	"travel/common/enum"
	"travel/common/xerr"

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
	if affected := l.svcCtx.DB.Model(&model.Content{}).Where("id = ?", req.Id).Scan(&content).RowsAffected; affected == 0 {
		return nil, errors.Wrap(xerr.NewErrMsg("该内容不存在"), "该内容不存在")
	}

	// 用户信息
	var userInfoView types.UserInfoView
	userId := content.UserId
	info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: userId})
	// todo：测试
	fmt.Println(info)
	_ = copier.Copy(&userInfoView, &info)
	content.UserInfo = userInfoView

	// 未登录
	if loginUserId == 0 {
		content.IsFavored = false
		content.IsLiked = false
	} else {
		// 是否点赞
		var isLiked bool
		l.svcCtx.DB.Model(&model.Like{}).Select("liked_status").Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.VIDEO, content.Id).Scan(&isLiked)
		content.IsLiked = isLiked

		// 是否收藏
		var favor model.Favor
		if err := l.svcCtx.DB.Model(&model.Favor{}).Where("user_id = ? and item_type = ? and item_id = ?", loginUserId, enum.VIDEO, content.Id).First(&favor).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				content.IsFavored = false
			}
		} else {
			content.IsFavored = true
		}
	}

	return &types.ContentDetailResp{
		ContentDetail: content,
	}, nil
}
