package rover

import (
	"github.com/leonardogazio/mars_rover_navigation/proto/pb"
)

// Position ...
type Position struct {
	X, Y int32
}

// Plateau ...
type Plateau struct {
	W, H int32
}

// Rover ...
type Rover struct {
	Position             Position
	Angle                int32
	SquadIncomingMessage string
}

// NewPlateau - Returns a new Plateau instance reference.
func NewPlateau() *Plateau {
	return &Plateau{W: 1000, H: 700}
}

// Rover - Returns a new Rover instance reference.
func NewRover() *Rover {
	return new(Rover)
}

// Contains - Checks if Plateau contains the Rover inside.
func (p *Plateau) Contains(r *Rover) bool {
	return r.Position.X >= 0 && r.Position.X <= p.W && r.Position.Y >= 0 && r.Position.Y <= p.H
}

// Rotate - Rotates the Rover object to selected direction L/R
func (r *Rover) Rotate(direction string) *Rover {
	if direction == "L" {
		r.Angle = r.Angle - 90
	} else {
		r.Angle = r.Angle + 90
	}
	r.Angle = (r.Angle + 360) % 360
	return r
}

// GoForward - Walks forward through the plateau based on current object orientation.
func (r *Rover) GoForward(p *Plateau, numberOfSteps int32) *Rover {
	switch r.Angle {
	case 90:
		r.Position.X += numberOfSteps
	case 180:
		r.Position.Y -= numberOfSteps
	case 270:
		r.Position.X -= numberOfSteps
	case 0:
		r.Position.Y += numberOfSteps
	}

	if r.SquadIncomingMessage == welcomeBackMessage {
		r.SquadIncomingMessage = ""
	}

	if !p.Contains(r) {
		r.SquadIncomingMessage = outsideOfPlateauMessage
	} else {
		if r.SquadIncomingMessage == outsideOfPlateauMessage {
			r.SquadIncomingMessage = welcomeBackMessage
		}
	}

	return r
}

func (r *Rover) ToProto() *pb.RoverPosition {
	return &pb.RoverPosition{
		X:                    r.Position.X,
		Y:                    r.Position.Y,
		Orientation:          pb.Orientation(r.Angle),
		SquadIncomingMessage: r.SquadIncomingMessage,
	}
}
