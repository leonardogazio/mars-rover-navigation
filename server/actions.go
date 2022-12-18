package server

import (
	"context"

	"github.com/leonardogazio/mars_rover_navigation/proto/pb"
	"github.com/leonardogazio/mars_rover_navigation/rover"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *MarsRoverNavigateServiceServer) SetPlateau(ctx context.Context, data *pb.SetPlateauReq) (*emptypb.Empty, error) {
	rover.PlateauInstance = rover.NewPlateau(data.GetUpper(), data.GetRight())
	return &emptypb.Empty{}, nil
}

func (s *MarsRoverNavigateServiceServer) LandDown(ctx context.Context, data *pb.LandDownReq) (*emptypb.Empty, error) {
	for _, r := range data.GetRoverList() {
		rover.RoverSquad[r.GetRoverID()] = &rover.Rover{
			ID: r.GetRoverID(),
			Position: rover.Position{
				X: r.GetPosition().GetX(),
				Y: r.GetPosition().GetY(),
			},
			Angle: int32(r.GetPosition().GetOrientation().Enum().Number()),
		}
	}
	return &emptypb.Empty{}, nil
}

func (s *MarsRoverNavigateServiceServer) Rotate(ctx context.Context, data *pb.RotateReq) (*pb.Rover, error) {
	r, found := rover.RoverSquad[data.GetRoverID()]
	if !found {
		return nil, roverNotFoundError
	}
	return r.Rotate(data.GetDirection().String()).ToProto(), nil
}

func (s *MarsRoverNavigateServiceServer) GoForward(ctx context.Context, data *pb.GoForwardReq) (*pb.Rover, error) {
	if rover.PlateauInstance == nil {
		return nil, plateauNotSet
	}

	r, found := rover.RoverSquad[data.GetRoverID()]
	if !found {
		return nil, roverNotFoundError
	}

	return r.GoForward(rover.PlateauInstance, data.GetNumberOfSteps()).ToProto(), nil
}

func (s *MarsRoverNavigateServiceServer) ShowSquad(ctx context.Context, data *emptypb.Empty) (*pb.RoverSquad, error) {
	squad := []*pb.Rover{}
	for _, r := range rover.RoverSquad {
		squad = append(squad, r.ToProto())
	}
	return &pb.RoverSquad{Rovers: squad}, nil
}
