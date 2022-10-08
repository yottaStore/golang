package hashbench

import (
	"crypto/rand"
	"github.com/zeebo/xxh3"
	"testing"
)

func benchmarkHashBlock64(size int, b *testing.B) {

	token := make([]byte, size*4096)
	_, err := rand.Read(token)
	if err != nil {
		b.Fatal("Error generating random token: ", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		xxh3.Hash(token)
	}

}

func benchmarkHashBlock128(size int, b *testing.B) {

	token := make([]byte, size*4096)
	_, err := rand.Read(token)
	if err != nil {
		b.Fatal("Error generating random token: ", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		xxh3.Hash128(token)
	}

}

func bHB64(size int, b *testing.B) {

	token := make([]byte, size)
	_, err := rand.Read(token)
	if err != nil {
		b.Fatal("Error generating random token: ", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		xxh3.Hash(token)
	}

}

func bHB128(size int, b *testing.B) {

	token := make([]byte, size)
	_, err := rand.Read(token)
	if err != nil {
		b.Fatal("Error generating random token: ", err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		xxh3.Hash128(token)
	}

}

func BenchmarkHashBlock1(b *testing.B)        { benchmarkHashBlock64(1, b) }
func BenchmarkHashBlock4104(b *testing.B)     { bHB64(4104, b) }
func BenchmarkHashBlock1_128(b *testing.B)    { benchmarkHashBlock128(1, b) }
func BenchmarkHashBlock4104_128(b *testing.B) { bHB128(4104, b) }

/*func BenchmarkHashBlock2_128(b *testing.B)  { benchmarkHashBlock128(2, b) }
func BenchmarkHashBlock4_128(b *testing.B)  { benchmarkHashBlock64(4, b) }
func BenchmarkHashBlock4(b *testing.B)      { benchmarkHashBlock64(4, b) }
func BenchmarkHashBlock8_128(b *testing.B)  { benchmarkHashBlock128(8, b) }
func BenchmarkHashBlock16_128(b *testing.B) { benchmarkHashBlock64(16, b) }

func BenchmarkHashBlock16(b *testing.B)     { benchmarkHashBlock64(16, b) }
func BenchmarkHashBlock32_128(b *testing.B) { benchmarkHashBlock128(32, b) }
func BenchmarkHashBlock64_128(b *testing.B) { benchmarkHashBlock128(64, b) }*/
