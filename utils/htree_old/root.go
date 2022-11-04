package htree_old

type Root struct {
	Root   *Node
	Prefix string
	Ops    []Update
	Hash   [16]byte
}

type Update struct {
	Type string
}

func (r *Root) Insert(n Node) error {

	if _, err := insert(&n, r.Root, true); err != nil {
		return err
	}
	return nil
}

func (r *Root) Update() error {

	return nil
}

func (r *Root) Delete() error {

	return nil
}

func (r *Root) Find() error {

	return nil
}

func (r *Root) Verify() error {

	return nil
}
