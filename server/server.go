package server

import (
	"fmt"
	"net"

	"github.com/leonardogazio/mars_rover_navigation/proto/pb"
	"github.com/leonardogazio/mars_rover_navigation/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// MarsRoverNavigateServiceServer ...
type MarsRoverNavigateServiceServer struct {
	pb.MarsRoverNavigateServiceServer
}

// StartGRPC ...
func StartGRPC() error {
	httpPort := utils.GetEnv("HTTP_PORT", "50000")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", httpPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	grpcService := grpc.NewServer()
	pb.RegisterMarsRoverNavigateServiceServer(grpcService, &MarsRoverNavigateServiceServer{})
	reflection.Register(grpcService)

	fmt.Printf("Application gRPC server initialized on port: %s\r\n", httpPort)

	return grpcService.Serve(listen)
}
