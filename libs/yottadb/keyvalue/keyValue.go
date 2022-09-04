package keyvalue

import (
	"libs/yottadb"
)

type Driver struct {
	Nodes *[]string
}

func (d Driver) Read(req yottadb.ReadRequest) (yottadb.ReadResponse, error) {

	var resp yottadb.ReadResponse

	return resp, nil
}

func (d Driver) Write(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	return resp, nil
}

func (d Driver) Update(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	return resp, nil
}

func (d Driver) Append(req yottadb.WriteRequest) (yottadb.WriteResponse, error) {

	var resp yottadb.WriteResponse

	return resp, nil
}

func (d Driver) Delete(req yottadb.WriteRequest) error {

	return nil
}

func New(nodes *[]string) (yottadb.Interface, error) {

	dbDriver := Driver{
		Nodes: nodes,
	}

	return dbDriver, nil

}
