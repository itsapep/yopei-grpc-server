package delivery

import (
	"log"
	"net"

	"github.com/itsapep/yopei-grpc/server/config"
	"github.com/itsapep/yopei-grpc/server/manager"
	"github.com/itsapep/yopei-grpc/server/service"
	"google.golang.org/grpc"
)

type YopeiGRPCServer struct {
	netListen      net.Listener
	server         *grpc.Server
	serviceManager manager.ServiceManager
}

func (ygs *YopeiGRPCServer) Run() {
	service.RegisterYopeiPaymentServer(ygs.server, ygs.serviceManager.YopeiService())
	log.Println("Server run", ygs.netListen.Addr().String())
	err := ygs.server.Serve(ygs.netListen)
	if err != nil {
		log.Fatalln("Failed to serve ...", err)
	}
}

func Server() *YopeiGRPCServer {
	yopeiGRPCServer := new(YopeiGRPCServer)
	c := config.NewConfig()
	listen, err := net.Listen("tcp", c.URL)
	if err != nil {
		log.Fatalln("Failed to listen ...", err)
	}
	grpcServer := grpc.NewServer()
	repoManager := manager.NewRepositoryManager()

	yopeiGRPCServer.serviceManager = manager.NewServiceManager(repoManager)
	yopeiGRPCServer.netListen = listen
	yopeiGRPCServer.server = grpcServer
	return yopeiGRPCServer
}
