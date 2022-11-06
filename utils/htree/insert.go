package htree

func insertShard(url string, weight uint64, root *Node, shouldUpdate bool) ([]*Node, error) {

	// TODO: use find
	// TODO: fail if parent is not a pnode

	coords, err := getCoords(url, root.Pointer, true)
	if err != nil {
		return nil, err
	}

	levels := len(coords)
	visitedNodes := make([]*Node, levels)
	visitedNodes[0] = root

	nodeType := T_VNODE
	parent := root
	children := root.Children
	for idx, coord := range coords[1:] {
		found := false
		insertionIdx := -1
		for cidx, child := range children {
			//insertionIdx = cidx

			if child.Pointer == coord {
				found = true
				visitedNodes[idx+1] = child
				child.Weight += weight
				parent = child
				children = child.Children
				break
			}
			// Make sure insertion respect the lexicographical order
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

			if len(parent.Children) == 0 {
				parent.Children = make([]*Node, 1)
			}

			if insertionIdx == -1 {
				//insertionIdx = len(parent.Children)
				parent.Children = append(children, &newNode)
			} else {
				parent.Children = append(children[:insertionIdx+1], children[insertionIdx:]...)
				parent.Children[insertionIdx] = &newNode
			}
			visitedNodes[idx+1] = &newNode
			parent = &newNode
			children = nil
		}

	}

	root.Weight += weight

	if shouldUpdate {

		for idx := levels - 1; idx > -1; idx-- {
			if err = visitedNodes[idx].Update(); err != nil {
				return visitedNodes, err
			}
		}

	}

	return visitedNodes, nil
}
