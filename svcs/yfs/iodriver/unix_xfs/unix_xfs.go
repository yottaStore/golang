package unix_xfs

import (
	"errors"
	"github.com/cornelk/hashmap"
	"github.com/yottaStore/golang/svcs/yfs/iodriver"
	"github.com/yottaStore/golang/utils/block"
	"github.com/yottaStore/golang/utils/utils"
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

	tb := block.Alloc(int(ts.Size) / block.BlockSize)
	if _, err := unix.Read(tfd, tb); err != nil {
		return nil, errors.New("error reading tails: " + err.Error())
	}

	tails, err := block.DeserializeTails(tb)
	if err != nil {
		return nil, errors.New("error deserializing tails: " + err.Error())
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
	bb := block.Alloc(int(ts.Blocks))
	if _, err := unix.Read(bfd, bb); err != nil {
		return nil, errors.New("error reading body: " + err.Error())
	}

	rb = append(rb, bb...)

	// Read appends
	// TODO: read appends blocks

	for _, tail := range tails {
		afd, err := unix.Open(string(tail.Pointer), ReadOpts, 0766)
		if err != nil {
			return nil, errors.New("error opening append: " + err.Error())
		}

		var appendStats unix.Stat_t
		if err := unix.Fstat(tfd, &appendStats); err != nil {
			return nil, errors.New("error getting append stat: " + err.Error())
		}

		appendBuffer := block.Alloc(int(tail.Length))
		if _, err := unix.Read(afd, appendBuffer); err != nil {
			return nil, errors.New("error reading append: " + err.Error())
		}

		rb = append(rb, appendBuffer...)
	}

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

func (d *IODriver) Compact(record string) error {

	return nil
}

func (d *IODriver) Merge(record string, payload []byte) error {

	return nil
}
