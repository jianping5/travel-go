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
	DataRpcConf zrpc.RpcClientConf
	DB          struct {
		DataSource string
	}
	COS struct {
		SecretID  string
		SecretKey string
		Region    string
		Bucket    string
	}
}
