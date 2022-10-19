package unix_xfs

import (
	"github.com/cornelk/hashmap"
	"github.com/yottaStore/golang/svcs/yfs/iodriver"
	"golang.org/x/sys/unix"
)

type IODriver struct {
	iodriver.Iodriver
	Namespace string
	Data      string
	Locks     *hashmap.Map[string, uint8]
}

const (
	DriverOpts = unix.O_DIRECT | unix.O_SYNC
	CreateOpts = DriverOpts | unix.O_WRONLY | unix.O_CREAT | unix.O_EXCL
	ReadOpts   = DriverOpts | unix.O_RDONLY
	AppendOpts = DriverOpts | unix.O_WRONLY | unix.O_APPEND
)

func (d *IODriver) Read(record string) ([]byte, error) {

	return read(record, d)
}

func (d *IODriver) Create(record string, payload []byte) error {

	return create(record, payload, d)
}

func (d *IODriver) Delete(record string) error {
	path := d.Data + "/" + record
	err := unix.Unlink(path + "/body")
	err = unix.Rmdir(path + "/append")
	err = unix.Unlink(path + "/tails")
	err = unix.Rmdir(path)
	return err
}

func (d *IODriver) Append(record string, payload []byte) error {

	return io_append(record, payload, d)
}

func (d *IODriver) Compact(record string) error {

	return compact(record, d)
}

func (d *IODriver) Merge(record string, payload []byte) error {

	return merge(record, d)
}
