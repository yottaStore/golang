package hTrie

import (
	"strings"
)

type Shard struct {
	Pointer string
	Weight  uint32
}

func (s *Shard) GetCoords(prefix string) []string {
	domain := strings.Split(s.Pointer, ":")[0]
	pointer := strings.Replace(domain, prefix, "", 1)
	coords := strings.Split(pointer, ".")
	coords[0] = s.Pointer
	levels := len(coords)

	for i, j := 0, levels-1; i < j; i, j = i+1, j-1 {
		coords[i], coords[j] = coords[j], coords[i]
	}
	coords[0] = prefix
	return coords
}
