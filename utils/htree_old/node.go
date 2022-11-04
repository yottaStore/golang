package htree_old

type NodeType uint8

const (
	T_SHARD NodeType = 1
	T_PNODE NodeType = 2
	T_VNODE NodeType = 3
	T_ROOT  NodeType = 4
)

const (
	HASH_SIZE        = 16
	NODE_FIELDS_SIZE = 9
	LEAF_FIELDS_SIZE = 5
)

type Node struct {
	Pointer  string
	Type     NodeType
	Weight   uint64
	Load     uint16
	Children []*Node
	Parent   *Node
	Hash     [HASH_SIZE]byte
}

type Insertion struct {
	Type    NodeType
	Pointer string
	Weight  uint64
	Load    uint16
}
