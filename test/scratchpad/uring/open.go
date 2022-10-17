package main

import (
	"github.com/ii64/gouring"
	"golang.org/x/sys/unix"
	"log"
	"unsafe"
)

func prepOpen(sqe *gouring.IoUringSqe, path string, flags int, fd, mode int) {
	gouring.PrepRW(gouring.IORING_OP_OPENAT, sqe, 0, unsafe.Pointer(&path), mode, 0)
	sqe.SetOpenFlags(uint32(flags))
}

func main() {
	h, err := gouring.New(256, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer h.Close()

	fileOpts := unix.O_RDONLY
	//mode := 0766
	path := "/tmp/test"

	sqe := h.GetSqe()
	prepOpen(sqe, path, 0766, 0, fileOpts)

	submitted, err := h.SubmitAndWait(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Submitted:", submitted)

	var cqe *gouring.IoUringCqe
	err = h.WaitCqe(&cqe)
	if err != nil {
		log.Fatal(err)
	} // check also EINTR

	log.Println("CQE: ", cqe)

}
