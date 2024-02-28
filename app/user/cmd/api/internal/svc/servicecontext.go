package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"travel/app/user/cmd/api/internal/config"
	"travel/app/user/cmd/rpc/user"
	"travel/common/initgorm"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
	DB      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initgorm.InitGorm(c.DB.DataSource)
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		DB:      db,
	}
}
