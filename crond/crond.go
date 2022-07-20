package crond

import (
	cron "github.com/robfig/cron/v3"
)

// NewWithSeconds returns a new Cron job runner, in the Local time zone.
// 返回一个支持至 秒 级别的 cron
func NewWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)

	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

// New returns a new Cron job runner, in the Local time zone.
// 返回一个支持至 分钟 级别的 cron
func New() *cron.Cron {
	return cron.New()
}
