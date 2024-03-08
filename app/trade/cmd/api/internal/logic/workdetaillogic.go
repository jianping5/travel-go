package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"travel/app/social/cmd/rpc/social"
	"travel/app/trade/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/xerr"

	"travel/app/trade/cmd/api/internal/svc"
	"travel/app/trade/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WorkDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWorkDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WorkDetailLogic {
	return &WorkDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WorkDetailLogic) WorkDetail(req *types.WorkDetailReq) (resp *types.WorkDetailResp, err error) {
	var work types.WorkView
	l.svcCtx.DB.Model(&model.Work{}).Where("id = ?", req.Id).Scan(&work)
	if work == (types.WorkView{}) {
		return nil, errors.Wrap(xerr.NewErrMsg("该商品不存在"), "该商品不存在")
	}

	var userInfoView types.UserInfoView
	userInfo, _ := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoReq{Id: work.UserId})
	_ = copier.Copy(userInfoView, userInfo)

	var copyright types.CopyrightView
	detail, err := l.svcCtx.SocialRpc.CopyrightDetail(l.ctx, &social.CopyrightDetailReq{Id: req.Id})
	_ = copier.Copy(&copyright, &detail)

	return &types.WorkDetailResp{
		Work:      work,
		UserInfo:  userInfoView,
		Copyright: copyright,
	}, nil
}
