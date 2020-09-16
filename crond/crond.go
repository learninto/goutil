package crond

import (
	"github.com/robfig/cron"
)

// New returns a new Cron job runner, in the Local time zone.
func New() *cron.Cron {
	return cron.New()
}
