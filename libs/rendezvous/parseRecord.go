package rendezvous

import (
	"errors"
	"log"
)

/*

Example of record string:

account@[driver:]collectionName/recordName[/recordRow[/subRow]]

*/

type ParsedRecord struct {
	Account    string
	Driver     string
	Collection string
	Record     string
}

func ParseRecord(record string) (ParsedRecord, error) {

	var parsed ParsedRecord

	atIndex := -1
	columnIndex := -1
	slashIndex := -1

Outer:
	for i, c := range record {

		switch c {
		case '@':
			if atIndex == -1 {
				atIndex = i
			} else {
				return parsed, errors.New("malformed record: too many @s")
			}

		case ':':
			columnIndex = i

		case '/':
			slashIndex = i
			log.Println("slash!")
			break Outer

		}

	}

	if columnIndex < 0 || atIndex < 0 || slashIndex < 0 {
		return parsed, errors.New("malformed record")
	}

	parsed = ParsedRecord{
		Account:    record[:atIndex],
		Driver:     record[atIndex+1 : columnIndex],
		Collection: record[columnIndex+1 : slashIndex],
		Record:     record[slashIndex:],
	}

	return parsed, nil

}
