package cmd

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"travel/app/social/cmd/api/internal/config"
	"travel/app/social/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var svcCtx *svc.ServiceContext

func Execute() {
	c := cron.New(cron.WithSeconds())

	ScheduleRun(c)

	fmt.Println("定时任务启动...")
	go c.Start()
	defer c.Stop()
	select {}
}

func init() {
	var c config.Config
	conf.MustLoad("app/social/cmd/api/etc/social.yaml", &c)
	svcCtx = svc.NewServiceContext(c)
}
