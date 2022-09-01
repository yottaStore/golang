package main

import (
	"fmt"
	"github.com/zeebo/xxh3"
	"time"
)

func main() {

	iters := 1000000
	start := time.Now()
	for i := 0; i < iters; i++ {

		tmp := xxh3.HashString128("Hello World")
		tmp = xxh3.HashString128("Hello World1")
		tmp = xxh3.HashString128("Hello World12")
		tmp = xxh3.HashString128("Hello World123")
		fmt.Println(tmp)

	}
	elapsedFull := time.Since(start)

	start = time.Now()
	hasher := xxh3.New()
	for i := 0; i < iters; i++ {
		hasher.WriteString("Hello World")
		tmp := hasher.Sum128()
		hasher.WriteString("1")
		tmp = hasher.Sum128()
		hasher.WriteString("2")
		tmp = hasher.Sum128()
		hasher.WriteString("3")
		tmp = hasher.Sum128()
		hasher.Reset()
		fmt.Println(tmp)
	}
	elapsedWrite := time.Since(start)

	fmt.Println(elapsedFull, elapsedWrite)

}
