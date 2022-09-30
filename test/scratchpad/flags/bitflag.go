package main

import (
	"log"
	"strconv"
)

type Bits uint64

const (
	F_DONE       Bits = 0b1
	F_CONTINUE        = 0b10
	F_COMPRESSED      = 0b100
)

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }

func main() {

	opts := F_DONE
	log.Println("Opts: ", strconv.FormatUint(uint64(opts), 2))

	check := opts & F_COMPRESSED
	test1 := Has(opts, F_DONE)

	test2 := Has(opts, F_CONTINUE)
	test3 := Has(opts, F_COMPRESSED)

	log.Println("Check:", strconv.FormatUint(uint64(check), 2))
	log.Println("Test 1: ", test1)
	log.Println("Test 2: ", test2)
	log.Println("Test 3: ", test3)

	log.Println("F_DONE:", F_DONE)
	log.Println("F_COMPRESSED:", F_COMPRESSED)

}
