package logic

import (
	"context"
	"github.com/pkg/errors"
	"travel/app/user/cmd/model"
	"travel/common/tool"
	"travel/common/xerr"

	"travel/app/user/cmd/rpc/internal/svc"
	"travel/app/user/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var UserAlreadyRegisterError = xerr.NewErrMsg("用户已注册")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 判断是否已注册
	affected := l.svcCtx.DB.Find(&model.User{}, "account = ?", in.Account).RowsAffected
	if affected != 0 {
		return nil, errors.Wrapf(UserAlreadyRegisterError, "用户 %s 已注册", in.Account)
	}

	user := &model.User{
		Account:  in.Account,
		Password: tool.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		Avatar:   "https://avatars.githubusercontent.com/u/83172782?v=4",
		Email:    "2712748478@qq.com",
	}
	l.svcCtx.DB.Create(user)

	return &pb.RegisterResp{
		Id:      user.Id,
		Account: user.Account,
		Email:   user.Email,
	}, nil
}
