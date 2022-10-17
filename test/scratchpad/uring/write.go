package main

import (
	"fmt"
	"github.com/ii64/gouring"
	"github.com/yottaStore/golang/libs/alloc"
	"golang.org/x/sys/unix"
	"log"
)

func main() {

	h, err := gouring.New(256, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer h.Close()

	path := "/loopfs/test"
	fileOpts := unix.O_RDWR | unix.O_CREAT | unix.O_DIRECT | unix.O_SYNC | unix.O_APPEND

	fd, err := unix.Open(path, fileOpts, 0677)
	block := alloc.New(2)
	copy(block, fmt.Sprintf("Hello world 1 \n", ))
	copy(block[4096:], fmt.Sprintf("Hello world 2 \n"))

	sqe := h.GetSqe()
	gouring.PrepWrite(sqe, fd, &block[0], len(block), 0)

	submitted, err := h.SubmitAndWait(1)
	if err != nil {
		log.Fatal("Error submitting: ", err)
	}
	log.Println("Submitted:", submitted)

	var cqe *gouring.IoUringCqe
	err = h.WaitCqe(&cqe)
	if err != nil {
		log.Fatal(err)
	} // check also EINTR

	log.Println("CQE: ", cqe)

	_ = cqe.UserData
	_ = cqe.Res
	_ = cqe.Flags
}