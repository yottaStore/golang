package block

import (
	"github.com/zeebo/assert"
	"testing"
)

func TestSerde(t *testing.T) {

	payload := []byte("hello wordl!")

	buff, err := Serialize(payload, 2, F_COMPRESSED)
	assert.NoError(t, err)

	b, err := Deserialize(buff)
	assert.NoError(t, err)

	assert.Equal(t, b[0].Body, payload)
	assert.Equal(t, b[0].Type, 2)
	assert.Equal(t, b[0].Flags, F_COMPRESSED)
	assert.Equal(t, b[0].Length, uint16(len(payload)))
	assert.Equal(t, b[0].Version, 0)
	assert.Equal(t, b[0].Reserved1, uint8(0))
	assert.Equal(t, b[0].Reserved2, uint8(0))

	// TODO: Add more tests
	// TODO: verify the hash

}
