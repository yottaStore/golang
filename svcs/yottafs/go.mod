module yottafs

go 1.19

require (
	iodrivers v0.0.1
	yottaclient v0.0.1
	yottaself v0.0.1

)

require golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect

replace (
	iodrivers v0.0.1 => ./../../libs/iodrivers
	yottaclient v0.0.1 => ./../../libs/yottaclient
	yottaself v0.0.1 => ./../yottaself
)
