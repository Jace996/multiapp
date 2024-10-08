package service

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/jace996/multiapp/dtm/api/dtm/v1"
	kitgrpc "github.com/jace996/multiapp/pkg/server/grpc"
	kithttp "github.com/jace996/multiapp/pkg/server/http"
)

func NewHttpServerRegister(msg *MsgService) kithttp.ServiceRegister {
	return kithttp.ServiceRegisterFunc(func(server *http.Server, middleware ...middleware.Middleware) {
		v1.RegisterMsgServiceHTTPServer(server, msg)
	})
}
func NewGrpcServerRegister(msg *MsgService) kitgrpc.ServiceRegister {
	return kitgrpc.ServiceRegisterFunc(func(server *grpc.Server, middleware ...middleware.Middleware) {
		v1.RegisterMsgServiceServer(server, msg)
	})
}
