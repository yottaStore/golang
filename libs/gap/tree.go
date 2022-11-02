package gap

import "log"

type Node struct {
	Pointer     string
	TotalWeight uint32
	Children    []*Node
}

type Shard struct {
	Weight uint16
	Coords []string
}

func (t *Node) Insert(s Shard) error {

	lastNode := t

	for _, coord := range s.Coords {

		var found bool
		for _, node := range lastNode.Children {
			if node.Pointer == coord {
				found = true
				node.TotalWeight += uint32(s.Weight)
				lastNode = node
				break
			}
		}

		if !found {
			newNode := Node{
				Pointer:     coord,
				TotalWeight: uint32(s.Weight),
			}

			lastNode.Children = append(lastNode.Children, &newNode)
			lastNode = &newNode
		}

	}

	return nil
}

func (t *Node) Print() {

	for _, child := range t.Children {
		log.Println("Child:", child)
		child.Print()
	}
}

func (t *Node) IsLeaf() bool {
	return len(t.Children) == 0
}
