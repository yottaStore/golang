module yottago

go 1.19

require (
yottafs v0.0.1
yottastore v0.0.1
	github.com/vmihailenco/msgpack/v5 v5.3.5
	github.com/zeebo/xxh3 v1.0.2
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261
)

require (
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
)

replace yottafs v0.0.1 => ./src/yottafs
replace yottastore v0.0.1 => ./src/yottastore

