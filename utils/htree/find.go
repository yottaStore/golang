package htree

import "errors"

func find(pointer string, root *Node, isShard bool) ([]*Node, error) {

	coords, err := getCoords(pointer, root.Pointer, isShard)
	if err != nil {
		return nil, err
	}

	levels := len(coords)
	visitedNodes := make([]*Node, levels)
	visitedNodes[0] = root

	children := root.Children
	for idx, coord := range coords[1:] {
		for _, child := range children {
			if child.Pointer == coord {
				visitedNodes[idx+1] = child
				children = child.Children
				break
			}
			if child.Pointer > coord {
				return visitedNodes, errors.New("node not found")
			}
		}
	}

	return visitedNodes, nil
}
