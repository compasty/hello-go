package functions

import (
	"testing"
)

func TestCommons(t *testing.T) {
	// Test the hypotenuse function.
	if hypot(3, 4) != 5 {
		t.Errorf("hypot(3, 4) != 5")
	}
}
