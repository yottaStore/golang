package main

import (
	"log"
	"scratchpad/size/size"
)

func main() {

	getSize, err := size.New(8, 8, 4096)

	if err != nil {
		log.Fatal("Error creating getSize: ", err)
	}

	testCase := func(buffLen, eCount, eRemainder int) {
		c, r := getSize(buffLen)

		if c != eCount {
			log.Fatal("Expected count not matched: ", buffLen, c, eCount)
		}

		if r != eRemainder {
			log.Fatal("Expected remainder not matched: ", buffLen, r, eRemainder)
		}

	}

	cases := [][3]int{
		{1, 1, 1},
		{4080, 1, 4080},
		{4085, 2, 4085},
		{4088, 2, 0},
		{4090, 2, 2},
		{8168, 2, 4080},
		{8168, 2, 4080},
		{8176, 3, 0},
	}

	for _, test := range cases {
		log.Println("Testing: ", test[0], test[1], test[2])
		testCase(test[0], test[1], test[2])

	}

}
