module yottafs

go 1.19

require (
	yottaclient v0.0.1
	yottanet v0.0.1
)

require golang.org/x/sys v0.0.0-20220829200755-d48e67d00261

replace (
	libs v0.0.1 => ./../../libs
	yottaclient v0.0.1 => ./../../libs/yottaclient
	yottanet v0.0.1 => ./../yottanet
)
