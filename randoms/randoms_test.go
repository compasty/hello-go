package randoms

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 10; i++ {
		a := rng.Intn(100)
		if a < 0 || a >= 100 {
			t.Errorf("Illegal random number: %d", a)
		}
	}
}
