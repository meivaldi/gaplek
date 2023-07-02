package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/meivaldi/gaplek/cmd"
	grpcHandler "github.com/meivaldi/gaplek/internal/delivery/grpc"
	pb "github.com/meivaldi/protobuf/gaplek"
	"google.golang.org/grpc"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	host := os.Getenv("APP_HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 50051
	}

	if host == "" {
		host = "0.0.0.0"
	}

	if port <= 0 {
		port = 50051
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("[GRPC] failed to listen %s:%d", host, port)
		return
	}

	grpcServer := grpc.NewServer()

	services := cmd.SetupService()
	gaplekHandler := grpcHandler.NewJitterHandler(services)

	pb.RegisterJitterServer(grpcServer, gaplekHandler)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	reflection.Register(grpcServer)

	log.Printf("[GRPC] serve at %s:%d", host, port)

	grpcServer.Serve(lis)
}
