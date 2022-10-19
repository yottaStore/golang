package unix_xfs

import (
	"errors"
	"github.com/yottaStore/golang/utils/block"
	"github.com/yottaStore/golang/utils/utils"
	"golang.org/x/sys/unix"
)

func compact(record string, d *IODriver) error {

	// Read body
	rb, err := read_body(record, d)
	if err != nil {
		return err

	}

	// Read tails
	tails, _, err := read_tails(record, d)
	if err != nil {
		return err
	}

	ab, err := read_appends(record, tails, d)
	if err != nil {
		return err
	}

	rb = append(rb, ab...)

	newBodyPath := d.Data + record + "/body." + utils.RandString(8)
	fd, err := unix.Open(newBodyPath, CreateOpts, 0766)
	defer close_fd(fd)
	if err != nil {
		return err
	}

	buff, err := block.Serialize(rb, block.BodyType, 0)
	if err != nil {
		return err
	}

	_, err = unix.Write(fd, buff)
	if err != nil {
		return err
	}

	// TODO: check for lock

	tailsPath := d.Data + record + "/tails"
	opts := AppendOpts
	tfd, err := unix.Open(tailsPath, opts, 0766)
	defer close_fd(tfd)
	if err != nil {
		return errors.New("error opening tails: " + err.Error())
	}

	err = unix.Rename(newBodyPath, d.Data+record+"/body")
	if err != nil {
		return err
	}

	// TODO: add fake block

	//err = unix.Fallocate(tfd, unix.FALLOC_FL_COLLAPSE_RANGE, 0, tsize)
	err = unix.Ftruncate(tfd, 0)
	if err != nil {
		return err
	}

	for _, tail := range tails {
		err = unix.Unlink(string(tail.Pointer))
		if err != nil {
			return err
		}
	}

	return nil
}
