package unix_xfs

import (
	"errors"
	"github.com/cornelk/hashmap"
	"github.com/yottaStore/golang/svcs/yfs/iodriver"
	"github.com/yottaStore/golang/utils/alloc"
	"github.com/yottaStore/golang/utils/utils"
	"golang.org/x/sys/unix"
	"log"
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

	var rb []byte

	// Read tails
	tailsPath := d.Data + "/" + record + "/tails"
	tfd, err := unix.Open(tailsPath, ReadOpts, 0766)
	if err != nil {
		return nil, errors.New("error opening tails: " + err.Error())
	}

	var ts unix.Stat_t
	if err := unix.Fstat(tfd, &ts); err != nil {
		return nil, errors.New("error getting tails stat: " + err.Error())
	}

	tb := alloc.New(int(ts.Blocks))
	if _, err := unix.Read(tfd, tb); err != nil {
		return nil, errors.New("error reading tails: " + err.Error())
	}

	// Read body
	bodyPath := d.Data + "/" + record + "/body"
	bfd, err := unix.Open(bodyPath, ReadOpts, 0766)
	if err != nil {
		return nil, errors.New("error opening body: " + err.Error())
	}

	var bs unix.Stat_t
	if err := unix.Fstat(tfd, &bs); err != nil {
		return nil, errors.New("error getting body stat: " + err.Error())
	}

	// TODO: check if size is correct
	bb := alloc.New(int(ts.Blocks))
	if _, err := unix.Read(bfd, bb); err != nil {
		return nil, errors.New("error reading body: " + err.Error())
	}

	rb = append(rb, bb...)
	rb = append(rb, tb...)

	// Read appends
	// TODO: read appends blocks

	return rb, nil
}

func (d *IODriver) Create(record string, payload []byte) error {

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

	_, err = unix.Write(fd, payload)

	return err
}

func (d *IODriver) Delete(record string) error {
	path := d.Data + "/" + record
	err := unix.Unlink(path)
	return err
}

func (d *IODriver) Append(record string, payload []byte) error {

	appendPath := d.Data + record + "/append/" + utils.RandString(8)

	fd, err := unix.Open(appendPath, CreateOpts, 0766)
	if err != nil {
		return err
	}

	_, err = unix.Write(fd, payload)
	if err != nil {
		return err
	}

	tailPath := d.Data + record + "/tails"
	td, err := unix.Open(tailPath, AppendOpts, 0766)
	if err != nil {
		log.Println("Crash here!: ", tailPath)
		return err
	}

	tailBlock := alloc.New(1)
	copy(tailBlock, appendPath)

	_, err = unix.Write(td, tailBlock)
	if err != nil {
		if err := unix.Unlink(appendPath); err != nil {
			return errors.New("error deleting failed append: " + err.Error())
		}
		return err
	}

	return nil
}

func (d *IODriver) Compact(record string) error {

	return nil
}

func (d *IODriver) Merge(record string, payload []byte) error {

	return nil
}
