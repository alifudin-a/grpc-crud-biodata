package main

import (
	"log"
	"net"

	grpcS "grpc-crud-biodata/grpc"
	pb "grpc-crud-biodata/proto/biodata"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":1180")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		log.Println("Connect GRPC to port :1180")
	}

	sr := grpcS.NewServer()

	grpcServer := grpc.NewServer()
	pb.RegisterBioServer(grpcServer, sr)
	grpcServer.Serve(lis)
}
