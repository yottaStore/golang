package yottadb

type DbDriverInterface interface {
	Read()
	Write()
	Update()
}

type DbDriver struct {
	DbDriverInterface
}
