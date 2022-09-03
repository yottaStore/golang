module yottafs

go 1.19

require (
	gossip v0.0.1
)

replace (
	gossip v0.0.1 => ./../pkgs/gossip
)