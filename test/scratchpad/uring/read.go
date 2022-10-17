package main

import (
	"github.com/ii64/gouring"
	"github.com/yottaStore/golang/utils/alloc"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	h, err := gouring.New(256, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer h.Close()

	fileOpts := unix.O_RDWR | unix.O_CREAT | unix.O_DIRECT | unix.O_SYNC | unix.O_APPEND
	fd, err := unix.Open("/tmp/test", fileOpts, 0766)

	sqe := h.GetSqe()
	b := alloc.New(3)
	gouring.PrepRead(sqe, fd, &b[0], len(b), 0)

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
	//log.Println("Buffer: ", b)
	log.Println("Buffer: ", string(b))

	_ = cqe.UserData
	_ = cqe.Res
	_ = cqe.Flags
}
