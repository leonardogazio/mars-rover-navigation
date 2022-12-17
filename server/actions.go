package server

import (
	"context"

	"github.com/leonardogazio/mars_rover_navigation/proto/pb"
)

func (s *MarsRoverNavigateServiceServer) Rotate(ctx context.Context, data *pb.RotateReq) (*pb.RoverPosition, error) {
	return RoverInstance.Rotate(data.GetDirection().String()).ToProto(), nil
}

func (s *MarsRoverNavigateServiceServer) GoForward(ctx context.Context, data *pb.GoForwardReq) (*pb.RoverPosition, error) {
	return RoverInstance.GoForward(PlateauInstance, data.GetNumberOfSteps()).ToProto(), nil
}
