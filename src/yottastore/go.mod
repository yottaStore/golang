module yottastore

go 1.19

require (
	libs v0.0.1
	yottanet v0.0.1

)

replace (
	libs v0.0.1 => ./../libs
	yottanet v0.0.1 => ./../yottanet

)
