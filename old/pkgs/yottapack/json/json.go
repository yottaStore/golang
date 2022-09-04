package json

import (
	"encoding/json"
	"yottaStore/yottaStore-go/src/pkgs/yottapack"
)

type Packer[T any] struct {
}

func (j Packer[T]) Pack(value T) ([]byte, error) {

	buff, err := json.Marshal(value)

	if err != nil {
		return []byte{}, err
	}

	return buff, nil
}

func (j Packer[T]) Unpack(buff []byte) (T, error) {

	var value T
	err := json.Unmarshal(buff, value)

	if err != nil {
		return value, err
	}

	return value, nil
}

func New[T any]() yottapack.Packer[T] {

	packer := Packer[T]{}

	return packer
}