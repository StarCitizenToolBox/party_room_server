package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"party_room_server/protos"
	"party_room_server/service"
)

func main() {
	StartGRPCService()
}

func StartGRPCService() {
	lis, err := net.Listen("tcp", "localhost:39399")
	if err != nil {
		panic("failed to listen port :" + err.Error())
	}
	grpcServer := grpc.NewServer()
	protos.RegisterPingServiceServer(grpcServer, &service.IndexServiceImpl{})
	fmt.Println("StartGRPCService ....")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
