package htree_old

func New(prefix string) Root {

	rootNode := Node{
		Pointer: prefix,
		Type:    T_ROOT,
	}

	r := Root{
		Prefix: prefix,
		Root:   &rootNode,
	}

	return r
}
