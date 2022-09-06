package rendezvous

import "github.com/zeebo/xxh3"

func NewHash(nodes []string) func(string) []byte {

	return func(key string) []byte {

		hashed := make([]byte, 0)

		hasher := xxh3.New()
		hasher.WriteString(key)
		bytes := hasher.Sum128().Bytes()
		hashed = append(hashed, bytes[:]...)
		for _, node := range nodes {
			hasher.WriteString(node)
			bytes := hasher.Sum128().Bytes()
			hashed = append(hashed, bytes[:]...)
		}

		return hashed
	}
}
