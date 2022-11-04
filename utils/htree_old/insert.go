package htree_old

import (
	"errors"
	"strings"
)

func getCoords(url, prefix string, isShard bool) []string {

	domain := strings.Split(url, ":")[0]
	pointer := strings.Replace(domain, prefix, "", 1)
	coords := strings.Split(pointer, ".")
	if isShard {
		coords[0] = url
	}
	levels := len(coords)

	for i, j := 0, levels-1; i < j; i, j = i+1, j-1 {
		coords[i], coords[j] = coords[j], coords[i]
	}
	coords[0] = prefix
	return coords
}

func insert(src, dst *Node, prefix string, shouldUpdate bool) ([]*Node, error) {

	switch src.Type {
	case T_SHARD:
		return insertShard(src, dst, prefix, shouldUpdate)
	case T_PNODE:
		return insertNode(src, dst, prefix, shouldUpdate)
	default:
		return nil, errors.New("invalid node type")
	}

}

func insertNode(src, dst *Node, prefix string, shouldUpdate bool) ([]*Node, error) {

	var parent *Node
	children := []*Node{dst}
	coords := getCoords(src.Pointer, prefix, false)
	levels := len(coords)
	visitedNodes := make([]*Node, levels)

	for idx, coord := range coords {

		var t NodeType
		if idx == levels-1 {
			t = T_PNODE
		} else {
			t = T_VNODE
		}

		found := false
		for cidx, child := range children {

			if child.Pointer > coord {
				found = true
				newNode := Node{
					Pointer: coord,
					Type:    t,
					Weight:  src.Weight,
					Load:    src.Load,
					Parent:  parent,
				}

				// Make sure insertion keep lexicographical order
				parent.Children = append(children[:cidx+1], children[cidx:]...)
				parent.Children[cidx] = &newNode
				visitedNodes[idx] = &newNode
				parent = &newNode
				children = nil
				break
			}

			if child.Pointer == coord {
				found = true
				child.Weight += src.Weight
				child.Load += src.Load
				visitedNodes[idx] = child
				parent = child
				children = child.Children
				break
			}
		}

		if !found {
			newNode := Node{
				Pointer: coord,
				Type:    t,
				Weight:  src.Weight,
				Load:    src.Load,
				Parent:  parent,
			}

			parent.Children = append(children, &newNode)
			visitedNodes[idx] = &newNode
			parent = &newNode
			children = nil
		}

	}

	return nil, nil
}

func insertShard(src, dst *Node, prefix string, shouldUpdate bool) ([]*Node, error) {

	return nil, nil
}
