module yottaStore

go 1.19

require (
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261
	yottanet v0.0.1
	yottaclient v0.0.1
)

replace (
yottanet v0.0.1 => ./../yottanet
yottaclient v0.0.1 => ./../yottaclient
)
