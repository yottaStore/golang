package htree

func insertShard(url string, weight uint64, root *Node, shouldUpdate bool) ([]*Node, error) {

	coords, err := getCoords(url, root.Pointer, false)
	if err != nil {
		return nil, err
	}

	levels := len(coords)
	visitedNodes := make([]*Node, levels)

	nodeType := T_VNODE
	parent := root
	children := root.Children
	for idx, coord := range coords {
		found := false
		insertionIdx := 0
		for cidx, child := range children {
			if child.Pointer == coord {
				found = true
				visitedNodes[idx] = child
				child.Weight += weight
				parent = child
				children = child.Children
				break
			}
			if child.Pointer > coord {
				insertionIdx = cidx
				break
			}
		}

		if !found {

			if idx == levels-1 {
				nodeType = T_SHARD
			}

			newNode := Node{
				Pointer: coord,
				Type:    nodeType,
				Weight:  weight,
				Parent:  parent,
			}

			// Make sure insertion keep the lexicographical order
			if insertionIdx == 0 {
				parent.Children = make([]*Node, 1)
			} else {
				parent.Children = append(children[:insertionIdx+1], children[insertionIdx:]...)
			}
			parent.Children[insertionIdx] = &newNode
			visitedNodes[idx] = &newNode
			parent = &newNode
			children = nil
		}

	}

	if shouldUpdate {

		for idx := levels - 1; idx >= 0; idx-- {
			if err = visitedNodes[idx].Update(); err != nil {
				return visitedNodes, err
			}
		}

		if err = root.Update(); err != nil {
			return visitedNodes, err
		}
	}

	return visitedNodes, nil
}
