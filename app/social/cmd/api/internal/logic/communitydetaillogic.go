package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"travel/app/social/cmd/model"
	"travel/app/user/cmd/rpc/pb/pb"
	"travel/common/ctxdata"
	"travel/common/tool"
	"travel/common/xerr"

	"travel/app/social/cmd/api/internal/svc"
	"travel/app/social/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommunityDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityDetailLogic {
	return &CommunityDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommunityDetailLogic) CommunityDetail(req *types.CommunityDetailReq) (resp *types.CommunityDetailResp, err error) {
	loginUserId := ctxdata.GetUidFromCtx(l.ctx)
	// 获取社区信息
	var community model.Community
	if affected := l.svcCtx.DB.Take(&model.Community{}, "id = ?", req.Id).Scan(&community).RowsAffected; affected == 0 {
		return nil, errors.Wrap(xerr.NewErrMsg("该社区不存在"), "该社区不存在")
	}
	var communityView types.CommunityView
	_ = copier.Copy(&communityView, &community)
	communityView.CreateTime = tool.TimeToString(community.CreateTime)

	// 获取用户信息
	userId := community.UserId
	userInfo, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &pb.UserInfoReq{Id: userId, LoginUserId: loginUserId})
	if err != nil {
		return nil, err
	}

	return &types.CommunityDetailResp{
		Community: communityView,
		UserId:    userId,
		Account:   userInfo.Account,
		Avatar:    userInfo.Avatar,
	}, nil
}
