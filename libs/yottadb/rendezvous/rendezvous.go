package rendezvous

import (
	"fmt"
	"github.com/zeebo/xxh3"
	"sort"
)

func Uint16(u xxh3.Uint128) []uint16 {
	return []uint16{
		uint16(u.Hi >> 0x30), uint16(u.Hi >> 0x20),
		uint16(u.Hi >> 0x10), uint16(u.Hi),
		uint16(u.Lo >> 0x30), uint16(u.Lo >> 0x20),
		uint16(u.Lo >> 0x10), uint16(u.Lo),
	}
}

type Finder struct {
	HashKey string
}

type NodeMap []string

func findNodes(record string, nodes NodeMap, f Finder, count int) (NodeMap, error) {

	hashedRecord := xxh3.HashString128(record + f.HashKey)
	tmpArray := NewNodeTree(Uint16(hashedRecord))

	fmt.Println(tmpArray)
	sort.Sort(tmpArray)
	fmt.Println(tmpArray)

	idxs := tmpArray.idxs[:count]
	pickedNodes := make(NodeMap, count)

	for i := 0; i < count; i++ {
		index := idxs[i] % len(nodes)
		pickedNodes[i] = nodes[index]
	}

	return pickedNodes, nil
}

func (f Finder) ParseRecord(record string) (ParsedRecord, error) {
	parsedRecord, err := ParseRecord(record)
	return parsedRecord, err
}

func (f Finder) FindNodes(record ParsedRecord, nodes NodeMap, count int) (NodeMap, error) {
	result, err := findNodes(record.TableIdentifier, nodes, f, count)
	return result, err
}

func (f Finder) FindShard(record ParsedRecord, nodes NodeMap, count int) (NodeMap, error) {
	result, err := findNodes(record.RecordIdentifier, nodes, f, count)
	return result, err
}
