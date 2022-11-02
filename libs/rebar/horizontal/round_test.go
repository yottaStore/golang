package horizontal

import (
	"github.com/zeebo/assert"
	"testing"
)

func TestRound(t *testing.T) {

	h := []byte{32, 23, 1, 27}
	w := []uint32{1, 1, 1, 1}

	i1 := Round(h, w)
	i2 := Round(h, nil)

	assert.Equal(t, i1, []int{0, 3, 1, 2})
	assert.Equal(t, i2, i1)

	h2 := []uint16{1, 2, 3, 4}
	w2 := []uint32{16, 2, 8, 3}

	i3 := Round(h2, w2)
	i4 := Round(h2, nil)

	assert.Equal(t, i3, []int{2, 0, 3, 1})
	assert.Equal(t, i4, []int{3, 2, 1, 0})

}
