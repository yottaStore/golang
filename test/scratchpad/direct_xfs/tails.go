package main

import (
	"encoding/binary"
	"github.com/fxamacker/cbor/v2"
	"github.com/yottaStore/golang/svcs/yfs/iodriver/unix_xfs"
	"github.com/yottaStore/golang/utils/block"
	"golang.org/x/sys/unix"
	"log"
)

func main() {

	tailsPath := "/tmp/yfs/data/test/tails"
	tfd, err := unix.Open(tailsPath, unix_xfs.ReadOpts, 0766)
	if err != nil {
		log.Fatal("error opening tails: " + err.Error())
	}

	var ts unix.Stat_t
	if err := unix.Fstat(tfd, &ts); err != nil {
		log.Fatal("error getting tails stat: " + err.Error())
	}

	count := int(ts.Size) / block.BlockSize

	tb := block.Alloc(count)

	if _, err := unix.Read(tfd, tb); err != nil {
		log.Fatal("error reading tails: " + err.Error())
	}

	b, err := block.Deserialize(tb)
	if err != nil {
		log.Fatal("error deserializing tails: " + err.Error())
	}

	for _, db := range b {

		var st block.SerializedTail

		err := cbor.Unmarshal(db.Body, &st)
		if err != nil {
			log.Fatal("error unmarshaling: " + err.Error())
		}

		t := block.Tail{
			Pointer: st[0],
			Length:  binary.BigEndian.Uint16(st[1]),
			Hash:    st[2],
		}

		log.Println("Pointer:", string(st[0]))
		log.Println("Length:", t.Length)
		log.Println("Hash:", string(st[2]))
	}

}
