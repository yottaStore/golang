package unix_xfs

import (
	"github.com/cornelk/hashmap"
	"github.com/yottaStore/golang/svcs/yfs/io_driver"
	"golang.org/x/sys/unix"
	"log"
)

func New(namespace string) (io_driver.IODriver, error) {

	var d IODriver

	data := namespace + "/data/"

	err := unix.Access(namespace, unix.O_RDWR)
	switch err {
	case nil:
	case unix.ENOENT:
		if err := unix.Mkdir(namespace, 0766); err != nil {
			log.Println("Error instantiating driver: ", err)
			return &d, err
		}
	default:
		return &d, err
	}

	err = unix.Access(data, unix.O_RDWR)
	switch err {
	case nil:
	case unix.ENOENT:
		if err := unix.Mkdir(data, 0766); err != nil {
			log.Println("Error instantiating driver: ", err)
			return &d, err
		}
	default:
		log.Println("Error instantiating driver: ", err)
		return &d, err
	}

	driver := IODriver{
		Namespace: namespace,
		Data:      data,
		Locks:     hashmap.New[string, uint8](),
	}

	return &driver, nil
}
