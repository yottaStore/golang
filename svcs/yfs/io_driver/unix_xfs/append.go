package unix_xfs

import (
	"errors"
	"github.com/yottaStore/golang/utils/block"
	"github.com/yottaStore/golang/utils/utils"
	"golang.org/x/sys/unix"
)

func io_append(record string, payload []byte, d *IODriver) error {
	appendPath := d.Data + record + "/append/" + utils.RandString(8)

	fd, err := unix.Open(appendPath, CreateOpts, 0766)
	if err != nil {
		return err
	}

	buff, err := block.Serialize(payload, block.AppendType, 0)
	if err != nil {
		return err
	}

	_, err = unix.Write(fd, buff)
	if err != nil {
		return err
	}

	tailPath := d.Data + record + "/tails"
	td, err := unix.Open(tailPath, AppendOpts, 0766)
	if err != nil {
		return err
	}

	t := block.Tail{
		Pointer: []byte(appendPath),
		Length:  1,
		Hash:    []byte("djjdjd"),
	}

	tails, err := block.SerializeTails([]block.Tail{t}, block.F_COMPRESSED)

	_, err = unix.Write(td, tails)
	if err != nil {
		if err := unix.Unlink(appendPath); err != nil {
			return errors.New("error deleting failed append: " + err.Error())
		}
		return err
	}

	return nil
}
