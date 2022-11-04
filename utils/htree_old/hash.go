package htree_old

import (
	"encoding/binary"
	"errors"
	"github.com/zeebo/xxh3"
)

func computeHash(n *Node) ([HASH_SIZE]byte, error) {

	switch n.Type {
	case T_SHARD:
		return shardHash(n)
	case T_VNODE | T_PNODE | T_ROOT:
		return nodeHash(n)
	default:
		return [16]byte{}, errors.New("invalid node type")
	}

}

func shardHash(n *Node) ([HASH_SIZE]byte, error) {

	size := len(n.Pointer) + 8 + 1
	buff := make([]byte, 0, size)
	buff = append(buff, n.Pointer...)
	buff = binary.BigEndian.AppendUint64(buff, n.Weight)
	buff[size-1] = byte(n.Type)

	return xxh3.Hash128(buff).Bytes(), nil
}

func nodeHash(n *Node) ([HASH_SIZE]byte, error) {

	size := len(n.Pointer) + HASH_SIZE*len(n.Children) + 8 + 4 + 1
	buff := make([]byte, 0, size)
	buff = append(buff, n.Pointer...)
	for _, child := range n.Children {
		buff = append(buff, child.Hash[:]...)
	}
	buff = binary.BigEndian.AppendUint64(buff, n.Weight)
	buff = binary.BigEndian.AppendUint16(buff, n.Load)
	buff[size-1] = byte(n.Type)

	return xxh3.Hash128(buff).Bytes(), nil

}
