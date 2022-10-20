package main

import (
	"github.com/ii64/gouring"
	"golang.org/x/sys/unix"
	"log"
)

func prepFallocate(sqe *gouring.IoUringSqe, fd, mode, offset, len int) {
	gouring.PrepRW(gouring.IORING_OP_FALLOCATE, sqe, fd, nil, mode, uint64(offset))
	sqe.IoUringSqe_Union2 = gouring.IoUringSqe_Union2(len)

}

func main() {
	h, err := gouring.New(256, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer h.Close()

	fileOpts := unix.O_RDWR | unix.O_DIRECT | unix.O_SYNC
	fd, err := unix.Open("/loopfs/test", fileOpts, 0766)

	sqe := h.GetSqe()
	prepFallocate(sqe, fd, unix.FALLOC_FL_COLLAPSE_RANGE, 0, 4096)

	submitted, err := h.SubmitAndWait(1)
	if err != nil {
		log.Fatal("Error submitting", err)
	}
	log.Println("Submitted:", submitted)

	var cqe *gouring.IoUringCqe
	err = h.WaitCqe(&cqe)
	if err != nil {
		log.Fatal(err)
	} // check also EINTR

	log.Println("CQE: ", cqe)
	//log.Println("Buffer: ", b)
	//log.Println("Buffer: ", string(b))

}
