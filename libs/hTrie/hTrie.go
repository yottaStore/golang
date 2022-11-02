package hTrie

type Trie struct {
	Root   *Node
	Prefix string
	Ops    []Update
	//Hash   [16]byte
}

func (t *Trie) Print() {
	t.Root.Print()
}

type Update struct {
	Type    string
	Pointer string
	Weight  uint32
}

func (t *Trie) Insert(s Shard) error {

	if err := t.Root.Insert(s, t.Prefix); err != nil {
		return err
	}

	op := Update{
		Type:    "insert",
		Pointer: s.Pointer,
		Weight:  s.Weight,
	}

	t.Ops = append(t.Ops[1:], op)

	return nil

}

func (t *Trie) Verify() error {

	children := []*Node{t.Root}

	for _, child := range children {
		if err := child.Verify(); err != nil {
			return err
		}
	}

	return nil
}
