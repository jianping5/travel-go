package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"travel/app/social/cmd/rpc/social"
	"travel/app/trade/cmd/api/internal/config"
	"travel/app/user/cmd/rpc/user"
	"travel/common/initgorm"
)

type ServiceContext struct {
	Config    config.Config
	DB        *gorm.DB
	UserRpc   user.User
	SocialRpc social.Social
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initgorm.InitGorm(c.DB.DataSource)
	return &ServiceContext{
		Config:    c,
		DB:        db,
		UserRpc:   user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		SocialRpc: social.NewSocial(zrpc.MustNewClient(c.SocialRpcConf)),
	}
}
