package main

import (
	"fmt"
	"log"
	"net"

	"github.com/meivaldi/gaplek/cmd"
	grpcHandler "github.com/meivaldi/gaplek/internal/delivery/grpc"
	pb "github.com/meivaldi/protobuf/gaplek"
	"google.golang.org/grpc"
)

func main() {
	host := "0.0.0.0"
	port := 50051

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("[GRPC] failed to listen %s:%d", host, port)
		return
	}

	grpcServer := grpc.NewServer()

	services := cmd.SetupService()
	gaplekHandler := grpcHandler.NewJitterHandler(services)

	pb.RegisterJitterServer(grpcServer, gaplekHandler)
	log.Printf("[GRPC] serve at %s:%d", host, port)

	grpcServer.Serve(lis)
}
