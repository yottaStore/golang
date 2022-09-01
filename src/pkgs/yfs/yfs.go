package yfs

import "yottaStore/yottaStore-go/src/libs/drivers"

type Namespace struct {
	Path string
}

type Driver interface {
	Read(path string) []byte
}

type YfsSetupOptions struct {
	Path string
}

func New[T drivers.IoDriver](opts YfsSetupOptions) (Namespace, error) {

	nspace := Namespace{
		Path: opts.Path,
	}

	return nspace, nil
}
