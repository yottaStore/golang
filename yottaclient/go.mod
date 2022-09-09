module yottaclient

require (
	yottadb v0.0.1
)

replace (
	yottadb v0.0.1 => ../svcs/yottadb
)

go 1.19
