package setup

import (
	"fmt"
	"frascati/config"
	"frascati/listener"
	"frascati/pbuf"
	"log"
	"net"

	"google.golang.org/grpc"
)

func SetupGrpc(app App) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GetListenerPort()))
	if err != nil {
		log.Fatalf("cannot initiate GRCP listener: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	pbuf.RegisterGreeterServer(grpcServer, listener.NewCobaListener(app.Logger))

	return grpcServer, lis
}
