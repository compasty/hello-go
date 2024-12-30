package sorts

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	a := []int{6, 3, 9, 8, 1}
	sort.Ints(a)
	assert.Equal(t, a[0], 1)
	b := []int{6, 3, 9, 8, 1}
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	assert.Equal(t, b[0], 9)
}
