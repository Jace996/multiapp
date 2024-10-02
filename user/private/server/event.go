package server

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/jace996/multiapp/event"
	"github.com/jace996/multiapp/event/trace"
	kitconf "github.com/jace996/multiapp/pkg/conf"
	"github.com/jace996/multiapp/pkg/dal"
	uow2 "github.com/jace996/uow"
	"github.com/defval/di"
)

func NewEventServer(
	c *kitconf.Data,
	conn dal.ConnName,
	logger klog.Logger,
	uowMgr uow2.Manager,
	handlers []event.ConsumerHandler,
	container *di.Container,
) *event.ConsumerFactoryServer {
	e := c.Endpoints.GetEventMergedDefault(string(conn))
	srv := event.NewConsumerFactoryServer(e, container)
	srv.Use(event.ConsumerRecover(event.WithLogger(logger)), trace.Receive(), event.Logging(logger), event.ConsumerUow(uowMgr))
	for _, handler := range handlers {
		srv.Append(handler)
	}
	return srv
}
