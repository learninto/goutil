package crond

import (
	cron "github.com/robfig/cron/v3"
)

// New returns a new Cron job runner, in the Local time zone.
func New() *cron.Cron {
	return cron.New()
}
