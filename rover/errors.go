package rover

import "errors"

var (
	errInconsistentPlateauData = errors.New("inconsistent plateau data")
	errInvalidPlateauData      = errors.New("plateau coordinates should be a valid number")
	errInconsistentRoverData   = errors.New("inconsistent rover data")
	errInvalidRoverData        = errors.New("rover coordinates should be a valid number")
	errInvalidCompassDirection = errors.New("invalid compass direction")
	errMovementCommand         = errors.New("invalid movement command string")
)
