package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"travel/app/user/cmd/model"
	"travel/app/user/cmd/rpc/internal/svc"
	"travel/app/user/cmd/rpc/pb/pb"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	var user *pb.UserInfoResp
	affected := l.svcCtx.DB.Take(&model.User{}, "id = ?", in.Id).Scan(&user).RowsAffected

	// 注入是否关注
	var id int64
	if l.svcCtx.DB.Model(&model.Follow{}).Select("id").Where("user_id = ? and follow_user_id = ?", in.LoginUserId, in.Id).Scan(&id); id == 0 {
		user.IsFollowed = false
	} else {
		user.IsFollowed = true
	}

	if affected == 0 {
		return nil, errors.Wrap(UserNoExistsError, "抱歉，该用户不存在")
	}

	return user, nil
}
