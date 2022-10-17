package alloc

import (
	"strconv"
)

func FormatToken(sec int64, nsec int64) []byte {

	return []byte(strconv.FormatInt(sec, 36) +
		strconv.FormatInt(nsec, 36))
}
