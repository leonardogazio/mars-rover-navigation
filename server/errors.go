package server

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	roverNotFoundError = status.Error(codes.NotFound, "rover not found")
	plateauNotSet      = status.Error(codes.FailedPrecondition, "plateau not set, set plateau first, before landing and positioning rovers.")
)
