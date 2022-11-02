package rebar

type LevelType uint8

const (
	Region     LevelType = 0
	Zone                 = 1
	DataCenter           = 2
	Room                 = 3
	Rack                 = 4
	Node                 = 5
	Shards               = 6
)

type Tree struct {
	Levels []Level
	Size   uint32
}

type Level struct {
	Type        LevelType
	TotalWeight uint16
	Leaves      []Leaf
}

type Leaf struct {
	Weight  uint16
	Pointer string
}

type Shard struct {
	Weight  uint16
	Pointer string
}

func NewTreeFromShards(nodes []Shard) Tree {
	return Tree{}
}

func NewTree() Tree {

	var t Tree
	t.Levels = make([]Level, 0, 7)

	zone := Leaf{Weight: 0, Pointer: "zone1."}
	zones := Level{
		Type:        Zone,
		TotalWeight: 0,
		Leaves:      []Leaf{zone},
	}
	t.Levels[0] = zones

	dc := Leaf{Weight: 0, Pointer: "dc1."}
	dcs := Level{
		Type:        DataCenter,
		TotalWeight: 0,
		Leaves:      []Leaf{dc},
	}
	t.Levels[1] = dcs

	rack := Leaf{Weight: 0, Pointer: "rack1."}
	racks := Level{
		Type:        DataCenter,
		TotalWeight: 0,
		Leaves:      []Leaf{rack},
	}
	t.Levels[2] = racks

	shard1 := Leaf{Weight: 1, Pointer: "shard1."}
	shard2 := Leaf{Weight: 1, Pointer: "shard2."}
	shard3 := Leaf{Weight: 1, Pointer: "shard3."}
	shard4 := Leaf{Weight: 1, Pointer: "shard4."}
	shard5 := Leaf{Weight: 1, Pointer: "shard5."}
	shard6 := Leaf{Weight: 1, Pointer: "shard6."}
	shard7 := Leaf{Weight: 1, Pointer: "shard7."}
	shard8 := Leaf{Weight: 1, Pointer: "shard8."}
	shard9 := Leaf{Weight: 1, Pointer: "shard9."}
	shards := Level{
		Type:        DataCenter,
		TotalWeight: 0,
		Leaves:      []Leaf{shard1, shard2, shard3, shard4, shard5, shard6, shard7, shard8, shard9},
	}
	t.Levels[3] = shards

	return t
}