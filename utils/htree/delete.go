package htree

import "errors"

func deleteNode(url string, root *Node, shouldUpdate bool) ([]*Node, error) {

	visitedNodes, err := find(url, root, true)
	if err != nil {
		return visitedNodes, err
	}
	didx := len(visitedNodes) - 1
	deletedNode := visitedNodes[didx]

	if deletedNode.Children != nil {
		return visitedNodes, errors.New("cannot delete, node has childrens")
	}

	for idx, node := range deletedNode.Parent.Children {
		if node == deletedNode {
			deletedNode.Parent.Children = append(deletedNode.Parent.Children[:idx], deletedNode.Parent.Children[idx+1:]...)
			break
		}
	}

	for i := didx - 1; i > -1; i-- {
		visitedNodes[i].Load -= deletedNode.Load
		visitedNodes[i].Weight -= deletedNode.Weight

		if shouldUpdate {
			err = visitedNodes[i].Update()
			if err != nil {
				return visitedNodes, err
			}
		}
	}

	return visitedNodes[:didx], nil
}
