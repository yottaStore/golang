package htree

import (
	"errors"
	"log"
)

type Root struct {
	Root   *Node
	Prefix string
	Ops    []Update
}

type UpdateType uint8

const (
	INSERT_SHARD UpdateType = 1
	UPDATE_LOAD  UpdateType = 2
	DELETE_NODE  UpdateType = 3
)

type Update struct {
	Type   UpdateType
	Url    string
	Weight uint64
	Load   uint32
}

func (r *Root) Find(url string, isShard bool) ([]*Node, error) {
	return find(url, r.Root, isShard)
}

func (r *Root) Update(us []Update) error {

	visits := make([][]*Node, len(us))

	for idx, u := range us {
		var visitedNodes []*Node
		var err error
		switch u.Type {
		case INSERT_SHARD:
			visitedNodes, err = insertShard(u.Url, u.Weight, r.Root, false)

		case UPDATE_LOAD:
			visitedNodes, err = updateLoad(u.Url, u.Load, r.Root, false)
		case DELETE_NODE:
			visitedNodes, err = deleteNode(u.Url, r.Root, false)
		default:
			return errors.New("unknown update type")
		}
		visits[idx] = visitedNodes
		r.Ops = append(r.Ops, u)
		if err != nil {
			return err
		}
	}

	updates, err := mergeUpdates(visits)
	if err != nil {
		return err
	}

	for i := len(updates) - 1; i > -1; i-- {
		for _, node := range updates[i] {
			if err = node.Update(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (r *Root) InsertOld(us []Update) error {

	for _, u := range us {
		if _, err := insertShard(u.Url, u.Weight, r.Root, true); err != nil {
			return err
		}
		r.Ops = append(r.Ops, u)
	}

	return nil
}

func (r *Root) UpdateLoadOld(us []Update) error {

	for _, u := range us {
		r.Ops = append(r.Ops, u)
	}

	return nil
}

func (r *Root) Delete(us []Update) error {

	for _, u := range us {
		r.Ops = append(r.Ops, u)
	}

	return nil
}

func (r *Root) Verify() error {

	return r.Root.Verify(true, true)
}

func (r *Root) Print() {

	log.Println("Printing tree:", r)
	r.Root.Print()
}
