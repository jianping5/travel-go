package svc

import (
	"gorm.io/gorm"
	"travel/app/data/cmd/api/internal/config"
	"travel/common/initgorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initgorm.InitGorm(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
