module yottafs

go 1.19

require (
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261
	gossip v0.0.1
	yottaclient v0.0.1
)

replace (
gossip v0.0.1 => ./../pkgs/gossip
yottaclient v0.0.1 => ./../yottaclient
)
