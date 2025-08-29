package main

import (
	"flag"
	"fmt"

	"example/pay/i-kun.vip/pb/pay"
	"example/pay/internal/config"
	orderpayServer "example/pay/internal/server/orderpay"
	payServer "example/pay/internal/server/pay"
	"example/pay/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/pay.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pay.RegisterPayServer(grpcServer, payServer.NewPayServer(ctx))
		pay.RegisterOrderPayServer(grpcServer, orderpayServer.NewOrderPayServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
