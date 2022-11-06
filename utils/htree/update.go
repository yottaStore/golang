package htree

func updateLoad(url string, load uint32, root *Node, shouldUpdate bool) ([]*Node, error) {

	visitedNodes, err := find(url, root, false)
	if err != nil {
		return visitedNodes, err
	}

	for i := len(visitedNodes) - 1; i > -1; i-- {
		visitedNodes[i].Load += load
		if shouldUpdate {
			err = visitedNodes[i].Update()
			if err != nil {
				return visitedNodes, err
			}
		}
	}

	visitedNodes[len(visitedNodes)-1].Type = T_PNODE

	return visitedNodes, nil
}

func mergeUpdates(visits [][]*Node) ([][]*Node, error) {

	var updates [][]*Node
	//updates := make([][]*Node, len(visits[0]))
	updateMap := make(map[*Node]bool)

	for _, visit := range visits {

		if len(visit) > len(updates) {
			updates = append(updates, make([][]*Node, len(visit)-len(updates))...)
		}

		for idx, node := range visit {
			_, ok := updateMap[node]
			if !ok {
				updateMap[node] = true
				updates[idx] = append(updates[idx], node)
			}
		}
	}

	return updates, nil
}
