package yfs

type Namespace struct {
	Path string
}

type Driver interface {
	Read(path string) []byte
}

type YfsSetupOptions struct {
	Path string
}

func New[T Driver](opts YfsSetupOptions) (Namespace, error) {

	nspace := Namespace{
		Path: opts.Path,
	}

	return nspace, nil
}