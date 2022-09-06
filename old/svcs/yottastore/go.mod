module store

go 1.19

require yottaself v0.0.1

require yottadb v0.0.1

require (
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
)

replace (
	yottadb v0.0.1 => ./../../libs/yottadb
	yottaself v0.0.1 => ./../yottaself
)
