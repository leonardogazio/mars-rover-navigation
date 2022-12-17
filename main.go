package main

import (
	"github.com/leonardogazio/mars_rover_navigation/rover"
	"github.com/leonardogazio/mars_rover_navigation/server"
)

func main() {
	server.PlateauInstance = rover.NewPlateau()
	server.RoverInstance = rover.NewRover()
	server.StartGRPC()
}
