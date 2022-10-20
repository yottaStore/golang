package unix_xfs

import (
	"errors"
	"github.com/yottaStore/golang/utils/block"
	"golang.org/x/sys/unix"
	"log"
)

func close_fd(fd int) {
	err := unix.Close(fd)
	if err != nil {
		log.Println("error closing file: " + err.Error())
	}
}

func read_appends(record string, tails []block.Tail, d *IODriver) ([]byte, error) {

	var rb []byte

	for _, tail := range tails {
		afd, err := unix.Open(string(tail.Pointer), ReadOpts, 0766)
		if err != nil {
			close_fd(afd)
			return nil, errors.New("error opening append: " + err.Error())
		}

		var appendStats unix.Stat_t
		if err := unix.Fstat(afd, &appendStats); err != nil {
			close_fd(afd)
			return nil, errors.New("error getting append stat: " + err.Error())
		}

		appendBuffer := block.Alloc(int(tail.Length))
		if _, err := unix.Read(afd, appendBuffer); err != nil {
			close_fd(afd)
			return nil, errors.New("error reading append: " + err.Error())
		}
		close_fd(afd)

		blocks, err := block.Deserialize(appendBuffer)
		if err != nil {
			return nil, errors.New("error deserializing append: " + err.Error())
		}

		for _, bl := range blocks {
			rb = append(rb, bl.Body[:bl.Length]...)
		}

	}

	return rb, nil
}

func read_tails(record string, d *IODriver) ([]block.Tail, int64, error) {
	tailsPath := d.Data + "/" + record + "/tails"
	tfd, err := unix.Open(tailsPath, ReadOpts, 0766)
	defer close_fd(tfd)
	if err != nil {
		return nil, 0, errors.New("error opening tails: " + err.Error())
	}

	var ts unix.Stat_t
	if err := unix.Fstat(tfd, &ts); err != nil {
		return nil, 0, errors.New("error getting tails stat: " + err.Error())
	}

	size := int(ts.Size) / block.BlockSize
	tb := block.Alloc(size)
	if _, err := unix.Read(tfd, tb); err != nil {
		return nil, 0, errors.New("error reading tails: " + err.Error())
	}

	tails, err := block.DeserializeTails(tb)
	if err != nil {
		return nil, 0, errors.New("error deserializing tails: " + err.Error())
	}

	return tails, ts.Size, nil
}

func read_body(record string, d *IODriver) ([]byte, error) {

	// Read body
	bodyPath := d.Data + "/" + record + "/body"
	bfd, err := unix.Open(bodyPath, ReadOpts, 0766)
	defer close_fd(bfd)
	if err != nil {
		return nil, errors.New("error opening body: " + err.Error())
	}

	var bs unix.Stat_t
	if err := unix.Fstat(bfd, &bs); err != nil {
		return nil, errors.New("error getting body stat: " + err.Error())
	}

	size := int(bs.Size) / block.BlockSize
	bb := block.Alloc(size)
	if _, err := unix.Read(bfd, bb); err != nil {
		return nil, errors.New("error reading body: " + err.Error())
	}

	var rb []byte
	blocks, err := block.Deserialize(bb)
	for _, bl := range blocks {
		rb = append(rb, bl.Body[:bl.Length]...)
	}

	return rb, nil
}

func read(record string, d *IODriver) ([]byte, error) {
	var rb []byte

	// Read body
	bb, err := read_body(record, d)
	if err != nil {
		return nil, err
	}
	rb = append(rb, bb...)

	// Read tails
	tails, _, err := read_tails(record, d)
	if err != nil {
		return nil, err
	}

	// Read appends
	ab, err := read_appends(record, tails, d)
	if err != nil {
		return nil, err
	}

	rb = append(rb, ab...)

	return rb, nil
}
