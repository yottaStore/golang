module yottadb

go 1.19

require (
	github.com/fxamacker/cbor/v2 v2.4.0
	rendezvous v0.0.1
	yottafs v0.0.1
)

require (
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
)

replace rendezvous v0.0.1 => ../../libs/rendezvous

replace yottafs v0.0.1 => ../yottafs
