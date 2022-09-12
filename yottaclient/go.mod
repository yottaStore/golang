module yottaclient

require (
	github.com/fxamacker/cbor/v2 v2.4.0
	yottadb v0.0.1
)

require github.com/x448/float16 v0.8.4 // indirect

replace yottadb v0.0.1 => ../svcs/yottadb

go 1.19
