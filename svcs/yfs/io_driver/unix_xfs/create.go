package unix_xfs

import (
	"errors"
	"github.com/yottaStore/golang/utils/block"
	"golang.org/x/sys/unix"
)

func create(record string, payload []byte, d *IODriver) error {

	path := d.Data + record
	if err := unix.Mkdir(path, 0766); err != nil {
		return errors.New("error creating record directory: " + err.Error())
	}
	if err := unix.Mkdir(path+"/append", 0766); err != nil {
		return errors.New("error creating record directory: " + err.Error())
	}

	_, err := unix.Open(path+"/tails", CreateOpts, 0766)
	if err != nil {
		return err
	}

	fd, err := unix.Open(path+"/body", CreateOpts, 0766)
	if err != nil {
		return err
	}

	buff, err := block.Serialize(payload, block.BodyType, 0)
	if err != nil {
		return err
	}

	_, err = unix.Write(fd, buff)

	return err

}
