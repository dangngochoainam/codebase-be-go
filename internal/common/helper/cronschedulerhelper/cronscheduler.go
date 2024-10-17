package cronschedulerhelper

import (
	"log"

	"github.com/robfig/cron/v3"
)

type Job struct {
	Spec string
	Cmd  func()
}

func NewCronSchedulerHelper(jobs []*Job) *cron.Cron {
	c := cron.New()
	if len(jobs) > 0 {
		for _, job := range jobs {
			c.AddFunc(job.Spec, job.Cmd)
		}
	}
	c.Start()
	log.Println("Cron scheduler initialized")
	return c
}
