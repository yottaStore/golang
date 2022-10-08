package hashbench

import (
	"crypto/rand"
	"github.com/zeebo/xxh3"
	"log"
)

func main() {

	iterations := 1000
	blocksize := 4096
	token := make([]byte, blocksize)
	_, err := rand.Read(token)
	if err != nil {
		log.Fatal("Error generating data block: ", err)
	}
	log.Println(token)

	ref := xxh3.Hash(token)

	for i := 0; i < iterations; i++ {

		h := xxh3.Hash(token)

		if h != ref {
			log.Fatal("Error hashing")
		}

	}

}
