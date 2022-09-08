package rendezvous

import (
	"github.com/zeebo/xxh3"
	"log"
	"sort"
)

// TODO: Handle replication factor and

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
}

func findNodes(record string, nodes NodeMap, f Finder, opts RendezvousOptions) (NodeMap, error) {

	hashedRecord := xxh3.HashString128(record + f.HashKey)
	tmpArray := NewNodeTree(Uint16(hashedRecord))

	log.Println(tmpArray)
	sort.Sort(tmpArray)
	log.Println(tmpArray)

	idxs := tmpArray.idxs[:opts.Sharding]
	pickedNodes := make(NodeMap, opts.Sharding)

	for i := 0; i < opts.Replication; i++ {
		index := idxs[i] % len(nodes)
		log.Println("Index: ", index)
		pickedNodes[i] = nodes[index]
	}

	return pickedNodes, nil
}

func (f Finder) ParseRecord(record string) (ParsedRecord, error) {
	parsedRecord, err := ParseRecord(record)
	return parsedRecord, err
}

func (f Finder) FindPool(record ParsedRecord, nodes NodeMap, count int) (NodeMap, error) {
	opts := RendezvousOptions{
		Sharding: count,
	}
	result, err := findNodes(record.TableIdentifier, nodes, f, opts)
	return result, err
}

func (f Finder) FindShard(record ParsedRecord, nodes NodeMap, count int) (NodeMap, error) {
	opts := RendezvousOptions{
		Sharding: count,
	}
	result, err := findNodes(record.RecordIdentifier, nodes, f, opts)
	return result, err
}

func (f Finder) FindRecord(record string, nodes NodeMap, opts RendezvousOptions) (NodeMap, NodeMap, ParsedRecord, error) {

	log.Println("Start find record")

	var parsedRecord ParsedRecord
	parsedRecord, err := ParseRecord(record)
	if err != nil {
		return nil, nil, parsedRecord, err
	}

	log.Println("Nodes:", nodes)

	pool, err := findNodes(parsedRecord.TableIdentifier, nodes, f, opts)
	if err != nil {
		return nil, nil, parsedRecord, err
	}

	log.Println("Pool:", pool)

	shards, err := findNodes(parsedRecord.RecordIdentifier, pool, f, opts)
	if err != nil {
		return nil, nil, parsedRecord, err
	}

	log.Println("Shards: ", shards)

	return shards, pool, parsedRecord, nil

}
