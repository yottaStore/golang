package direct

import (
	"golang.org/x/sys/unix"
	"log"
)

func New(namespace string) (Driver, error) {

	var d Driver

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

	err = unix.Access(namespace+"/data/", unix.O_RDWR)
	switch err {
	case nil:
	case unix.ENOENT:
		if err := unix.Mkdir(namespace+"/data/", 0766); err != nil {
			log.Println("Error instantiating driver: ", err)
			return d, err
		}
	default:
		log.Println("Error instantiating driver: ", err)
		return d, err
	}

	driver := Driver{
		namespace: namespace,
	}

	return driver, nil
}
