package rendezvous

import (
	"errors"
	"github.com/zeebo/xxh3"
	"sort"
)

type ReplicationMode int64

const (
	Regional ReplicationMode = iota
	DualRegion
	MultiRegion
)

type RendezvousOptions struct {
	Replication int
	Mode        ReplicationMode
	Sharding    int
	HashKey     string
}

type NodeMap []string

func FindPool(record ParsedRecord, nodes NodeMap, opts RendezvousOptions) (
	NodeMap, error) {

	if opts.HashKey == "" {
		return NodeMap{}, errors.New("Hash key can't be null")
	}

	return findNodes(record.CollectionIdentifier, nodes, opts)
}

func FindNode(record ParsedRecord, pool NodeMap, opts RendezvousOptions) (
	NodeMap, error) {

	if opts.HashKey == "" {
		return NodeMap{}, errors.New("Hash key can't be null")
	}

	return findNodes(record.RecordIdentifier, pool, opts)
}

func FindRecord(record string, nodes NodeMap, opts RendezvousOptions) (NodeMap, NodeMap, ParsedRecord, error) {

	if opts.HashKey == "" {
		return NodeMap{}, NodeMap{}, ParsedRecord{}, errors.New("Hash key can't be null")
	}

	parsedRecord, err := ParseRecord(record)
	if err != nil {
		return nil, nil, parsedRecord, err
	}

	pool, err := findNodes(parsedRecord.CollectionIdentifier, nodes, opts)
	if err != nil {
		return nil, nil, parsedRecord, err
	}

	shards, err := findNodes(parsedRecord.RecordIdentifier, pool, opts)
	if err != nil {
		return nil, nil, parsedRecord, err
	}

	return shards, pool, parsedRecord, nil

}

func findNodes(record string, nodes NodeMap, opts RendezvousOptions) (NodeMap, error) {

	hashedRecord := xxh3.HashString128(record + opts.HashKey)
	tmpArray := NewNodeTree(Uint16(hashedRecord))
	sort.Sort(tmpArray)

	idxs := tmpArray.idxs[:opts.Sharding]
	pickedNodes := make(NodeMap, opts.Sharding)

	for i := 0; i < opts.Replication; i++ {
		index := idxs[i] % len(nodes)
		pickedNodes[i] = nodes[index]
		pickedNodes[i] = nodes[index]
	}

	return pickedNodes, nil
}

func Uint16(u xxh3.Uint128) []uint16 {
	return []uint16{
		uint16(u.Hi >> 0x30), uint16(u.Hi >> 0x20),
		uint16(u.Hi >> 0x10), uint16(u.Hi),
		uint16(u.Lo >> 0x30), uint16(u.Lo >> 0x20),
		uint16(u.Lo >> 0x10), uint16(u.Lo),
	}
}
