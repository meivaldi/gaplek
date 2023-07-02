package main

import (
	"log"
	"net"
	"os"

	"github.com/meivaldi/gaplek/cmd"
	grpcHandler "github.com/meivaldi/gaplek/internal/delivery/grpc"
	pb "github.com/meivaldi/protobuf/gaplek"
	"google.golang.org/grpc"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("[GRPC] failed to listen to port %s", port)
		return
	}

	grpcServer := grpc.NewServer()

	services := cmd.SetupService()
	gaplekHandler := grpcHandler.NewJitterHandler(services)

	pb.RegisterJitterServer(grpcServer, gaplekHandler)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	reflection.Register(grpcServer)

	log.Printf("[GRPC] serve at port %s", port)

	grpcServer.Serve(lis)
}
