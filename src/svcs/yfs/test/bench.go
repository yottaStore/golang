package main

import (
	"fmt"
	"time"
	"yottaStore/yottaStore-go/src/svcs/yfs/test/utils"
)

func main() {

	iters := 1000

	start := time.Now()
	for i := 0; i < iters; i++ {
		file := make([]byte, 4096*2)
		a := utils.Alignment(file, utils.AlignSize)
		offset := 0
		if a != 0 {
			offset = utils.AlignSize - a
		}
		file = file[offset : offset+utils.BlockSize]
		fmt.Println(file[0])
	}
	elapsedAlign := time.Since(start)

	startPrnt := time.Now()
	for i := 0; i < iters; i++ {
		file := make([]byte, 4096)
		fmt.Println(&file[0])
	}

	elapsedPrnt := time.Since(startPrnt)

	start = time.Now()

	fmt.Println(elapsedAlign, elapsedPrnt)

}
