package htree

import (
	"errors"
	"strings"
)

func getCoords(url, prefix string, isShard bool) ([]string, error) {

	domain := strings.Split(url, ":")[0]
	pidx := strings.Index(domain, prefix)
	if pidx == -1 {
		return nil, errors.New("wrong prefix")
	}
	coords := strings.Split(domain[:pidx], ".")
	if isShard {
		coords[0] = url
	}
	levels := len(coords)
	//coords = coords[0 : levels-1]
	//levels--

	for i, j := 0, levels-1; i < j; i, j = i+1, j-1 {
		coords[i], coords[j] = coords[j], coords[i]
	}
	coords[0] = prefix
	return coords, nil
}
