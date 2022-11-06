package htree

func New(prefix string) (*Root, error) {

	root := &Node{
		Pointer: prefix,
		Type:    T_ROOT,
	}

	tree := &Root{
		Root:   root,
		Prefix: prefix,
	}

	return tree, nil
}
