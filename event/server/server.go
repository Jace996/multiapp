package server

import (
	v12 "github.com/jace996/multiapp/event/api/v1"
	"github.com/jace996/multiapp/event/service"
	kitdi "github.com/jace996/multiapp/pkg/di"
	"github.com/goava/di"
)

var EventProviderSet = kitdi.NewSet(
	kitdi.NewProvider(service.NewEventService, di.As(new(v12.EventServiceServer))),
	service.NewGrpcServerRegister, service.NewHttpServerRegister,
)
