package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Salt        string
	UserRpcConf zrpc.RpcClientConf
	DB          struct {
		DataSource string
	}
}
