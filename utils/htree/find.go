package htree

import "errors"

func find(pointer string, root *Node) ([]*Node, error) {

	coords, err := getCoords(pointer, root.Pointer, false)
	if err != nil {
		return nil, err
	}

	levels := len(coords)
	visitedNodes := make([]*Node, 0, levels)

	for idx, coord := range coords {
		for _, child := range root.Children {
			if child.Pointer == coord {
				visitedNodes[idx] = child
				root = child
				break
			}
			if child.Pointer > coord {
				return visitedNodes, errors.New("node not found")
			}
		}
	}

	return visitedNodes, nil
}