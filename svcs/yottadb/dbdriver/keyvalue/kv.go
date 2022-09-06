package keyvalue

import "yottadb/dbdriver"

type Driver struct {
	dbdriver.Interface
}

func New() (dbdriver.Interface, error) {

	var d Driver

	return d, nil

}

func (d Driver) Read() {

}

func (d Driver) Write() {

}

func (d Driver) Delete() {

}
