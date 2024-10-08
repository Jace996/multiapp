package server

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/jace996/multiapp/pkg/job"
	"github.com/jace996/multiapp/sys/private/biz"
	"github.com/hibiken/asynq"
)

func NewJobServer(opt asynq.RedisConnOpt, log klog.Logger, handlers []*job.Handler) *job.Server {
	srv := job.NewServer(opt, job.WithQueues(map[string]int{
		string(biz.ConnName): 1,
	}))
	srv.Use(job.TracingServer(), job.Logging(log))
	job.RegisterHandlers(srv, handlers...)
	return srv
}
