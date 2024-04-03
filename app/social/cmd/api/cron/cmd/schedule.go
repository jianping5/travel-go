package cmd

import (
	"github.com/robfig/cron/v3"
	"travel/app/social/cmd/api/cron/cronx"
)

func ScheduleRun(c *cron.Cron) {
	_, _ = c.AddFunc(cronx.Daily(string(rune(1))), UpdateUserTag)
}
