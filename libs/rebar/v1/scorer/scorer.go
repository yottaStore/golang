package scorer

import (
	"errors"
	"github.com/yottaStore/golang/utils/htree"
)

var ErrNotEnoughChildren = errors.New("not enough children")

type Interface interface {
	Compute(nodes []*htree.Node, hash []byte, shards int) ([]*htree.Node, error)
}
