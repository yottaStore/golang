module store

go 1.19

require (
	yottadb v0.0.1
	yottaself v0.0.1
)

replace (
	yottadb v0.0.1 => ./../../libs/yottadb
	yottaself v0.0.1 => ./../yottaself
)
