package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"party_room_server/db"
	"party_room_server/protos"
	"party_room_server/service"
)

func main() {
	err := db.Connect()
	if err != nil {
		panic(err)
	}
	StartGRPCService()
}

func StartGRPCService() {
	lis, err := net.Listen("tcp", "localhost:39399")
	if err != nil {
		panic("failed to listen port :" + err.Error())
	}
	grpcServer := grpc.NewServer()
	protos.RegisterIndexServiceServer(grpcServer, &service.IndexServiceImpl{})
	fmt.Println("StartGRPCService ....")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
