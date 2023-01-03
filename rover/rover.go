package rover

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/leonardogazio/mars-rover-navigation/proto/pb"
)

// Position ...
type Position struct {
	X, Y int32
}

// Plateau ...
type Plateau struct {
	R, U int32
}

// Rover ...
type Rover struct {
	ID                   string
	Position             Position
	Angle                int32
	SquadIncomingMessage string
}

var (
	// RoverSquad ...
	RoverSquad map[string]*Rover = map[string]*Rover{}
	// PlateauInstance ...
	PlateauInstance *Plateau
)

// NewPlateau - Returns a new Plateau instance reference.
func NewPlateau(upper, right int32) *Plateau {
	return &Plateau{R: right, U: upper}
}

// Contains - Checks if Plateau contains the Rover inside.
func (p *Plateau) Contains(r *Rover) bool {
	return r.Position.X >= 0 && r.Position.X <= p.R && r.Position.Y >= 0 && r.Position.Y <= p.U
}

// Rotate - Rotates the Rover object to selected direction L/R
func (r *Rover) Rotate(direction string) *Rover {
	if direction == "L" {
		r.Angle -= 90
	} else {
		r.Angle += 90
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

func (r *Rover) ToProto() *pb.Rover {
	return &pb.Rover{
		RoverID: r.ID,
		Position: &pb.RoverPosition{
			X:                    r.Position.X,
			Y:                    r.Position.Y,
			Orientation:          pb.Orientation(r.Angle),
			SquadIncomingMessage: r.SquadIncomingMessage,
		},
	}
}

// ParseFile - Loads plateau and rover list and move them from file
func ParseFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	s.Scan()

	plateauCoords := strings.Split(s.Text(), " ")

	if len(plateauCoords) != 2 {
		return "", errInconsistentPlateauData
	}

	plateauX, err := strconv.Atoi(plateauCoords[0])
	if err != nil {
		return "", errInvalidPlateauData
	}

	plateauY, err := strconv.Atoi(plateauCoords[1])
	if err != nil {
		return "", errInvalidPlateauData
	}

	PlateauInstance = NewPlateau(int32(plateauY), int32(plateauX))

	res := ""

	for s.Scan() {
		roverData := strings.Split(s.Text(), " ")

		if len(roverData) != 3 {
			return "", errInconsistentRoverData
		}

		roverX, err := strconv.Atoi(roverData[0])
		if err != nil {
			return "", errInvalidRoverData
		}

		roverY, err := strconv.Atoi(roverData[1])
		if err != nil {
			return "", errInvalidRoverData
		}

		regx, _ := regexp.Compile("(^[NESW]$)")
		if !regx.MatchString(roverData[2]) {
			return "", errInvalidCompassDirection
		}

		roverObj := &Rover{
			Position: Position{X: int32(roverX), Y: int32(roverY)},
			Angle:    pb.Orientation_value[roverData[2]],
		}

		s.Scan()

		regx, _ = regexp.Compile(`^[LRM\.\,\+\-]*$`)
		if !regx.MatchString(s.Text()) {
			return "", errMovementCommand
		}

		for _, c := range s.Text() {
			switch cmd := fmt.Sprintf("%c", c); {
			case cmd == "L", cmd == "R":
				roverObj.Rotate(cmd)
			case cmd == "M":
				roverObj.GoForward(PlateauInstance, 1)
			}
		}

		res += fmt.Sprintf("%d %d %s\r\n", roverObj.Position.X, roverObj.Position.Y, pb.Orientation_name[roverObj.Angle])
	}

	return res, nil
}
