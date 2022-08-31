package main

import (
	"fmt"
	"github.com/zeebo/xxh3"
)

func main() {

	hasher := xxh3.New()

	hasher.Write([]byte("Hello world"))

	fmt.Println(hasher.Sum128())

	hasher.Write([]byte(" from yotta!"))

	fmt.Println(hasher.Sum128().Bytes())

	hasher.Reset()
	hasher.Write([]byte("Hello world from yotta!"))
	fmt.Println(hasher.Sum128().Bytes())

}
