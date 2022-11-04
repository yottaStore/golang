package hTrie

import (
	"github.com/zeebo/assert"
	"testing"
)

func TestShard(t *testing.T) {
	shard := Shard{
		Pointer: "hello.world.prefix.com:8080",
		Weight:  1,
	}

	coords := shard.GetCoords("prefix.com")

	tCoords := []string{"prefix.com", "world", "hello.world.prefix.com:8080"}
	assert.Equal(t, coords, tCoords)

}
