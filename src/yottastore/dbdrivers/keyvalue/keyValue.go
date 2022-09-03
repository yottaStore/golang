package keyvalue

import "yottastore/dbdrivers"

func New(nodes *[]string) (dbdrivers.Interface, error) {

	var dbDriver dbdrivers.Interface

	return dbDriver, nil

}
