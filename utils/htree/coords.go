package htree

import "strings"

func getCoords(url, prefix string, isShard bool) ([]string, error) {

	domain := strings.Split(url, ":")[0]
	pointer := strings.Replace(domain, prefix, "", 1)
	coords := strings.Split(pointer, ".")
	if isShard {
		coords[0] = url
	}
	levels := len(coords)

	for i, j := 0, levels-1; i < j; i, j = i+1, j-1 {
		coords[i], coords[j] = coords[j], coords[i]
	}
	coords[0] = prefix
	return coords, nil
}
