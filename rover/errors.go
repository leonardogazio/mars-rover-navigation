package rover

import "errors"

var (
	errInconsistentPlateauData = errors.New("inconsistent plateau data")
	errInconsistentRoverData   = errors.New("inconsistent rover data")
	errInvalidCompassDirection = errors.New("invalid compass direction")
	errMovementCommand         = errors.New("invalid movement command string")
)
