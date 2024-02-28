package auth

import (
	"context"
	"github.com/pkg/errors"
	"travel/app/user/cmd/model"
	"travel/app/user/cmd/rpc/user"
	"travel/common/tool"
	"travel/common/xerr"

	"travel/app/user/cmd/rpc/internal/svc"
	"travel/app/user/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var UserNoExistsError = xerr.NewErrMsg("用户不存在")
var UserPwdError = xerr.NewErrMsg("用户密码错误")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	var userId int64
	var err error
	userId, err = l.loginByAccount(in.Account, in.Password)
	if err != nil {
		return nil, err
	}

	// 生成 token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}

	return &user.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
	}, nil
}

func (l *LoginLogic) loginByAccount(account, password string) (int64, error) {
	var user model.User
	affected := l.svcCtx.DB.Take(&model.User{}, "account = ?", account).Scan(&user).RowsAffected
	if affected == 0 {
		return 0, errors.Wrapf(UserNoExistsError, "account:%s", account)
	}
	if tool.PasswordEncrypt(l.svcCtx.Config.Salt, password) != user.Password {
		return 0, errors.Wrap(UserPwdError, "用户密码错误")
	}

	return user.Id, nil
}
