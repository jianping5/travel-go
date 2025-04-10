package logic

import (
	"context"
	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type CopyrightListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCopyrightListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CopyrightListLogic {
	return &CopyrightListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CopyrightListLogic) CopyrightList(req *types.CopyrightListReq) (resp *types.CopyrightListResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	userId := req.UserId
	// 若传入的 userId 为 0，则表示为当前用户
	if userId == 0 {
		userId = ctxdata.GetUidFromCtx(l.ctx)
	}
	var copyrights []types.CopyrightView
	l.svcCtx.DB.Model(&model.Copyright{}).Where("user_id = ?", userId).Scan(&copyrights)

	for i, c := range copyrights {
		var content model.Content
		l.svcCtx.DB.Model(&model.Content{}).Select("title", "cover_url", "content").Where("id = ?", c.ItemId).Scan(&content)
		copyrights[i].Title = content.Title
		copyrights[i].CoverUrl = content.CoverUrl
		copyrights[i].Content = content.Content

		// 用户信息
		info, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: c.UserId, LoginUserId: loginUserId})
		copyrights[i].Account = info.Account
		copyrights[i].Avatar = info.Avatar
	}

	return &types.CopyrightListResp{
		List: copyrights,
	}, nil
}
