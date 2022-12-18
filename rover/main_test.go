package rover

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	PlateauInstance = NewPlateau(700, 1000)
	os.Exit(m.Run())
}
