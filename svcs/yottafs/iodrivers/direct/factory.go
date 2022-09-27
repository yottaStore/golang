package direct

import (
	"golang.org/x/sys/unix"
	"log"
)

func New(namespace string) (Driver, error) {

	var d Driver
	data := namespace + "/data/"

	err := unix.Access(namespace, unix.O_RDWR)
	switch err {
	case nil:
	case unix.ENOENT:
		if err := unix.Mkdir(namespace, 0766); err != nil {
			log.Println("Error instantiating driver: ", err)
			return d, err
		}
	default:
		return d, err
	}

	err = unix.Access(data, unix.O_RDWR)
	switch err {
	case nil:
	case unix.ENOENT:
		if err := unix.Mkdir(data, 0766); err != nil {
			log.Println("Error instantiating driver: ", err)
			return d, err
		}
	default:
		log.Println("Error instantiating driver: ", err)
		return d, err
	}

	driver := Driver{
		namespace: namespace,
		dataspace: data,
	}

	return driver, nil
}
