package record

import "errors"

type Record struct {
	Account      string
	Driver       string
	Collection   string
	Record       string
	PoolPointer  string
	ShardPointer string
}

// Example record: "account@driver:collection/record/subrecord"

func Parse(record string, isAdmin bool) (Record, error) {

	atIndex := -1
	colonIndex := -1
	firstSlashIndex := -1

	var r Record

	for i, c := range record {
		switch c {
		case '@':
			if atIndex != -1 {
				return r, errors.New("invalid record: multiple @")
			}
			atIndex = i
		case ':':
			if colonIndex != -1 {
				return r, errors.New("invalid record: multiple :")
			}
			colonIndex = i
		case '/':
			if firstSlashIndex == -1 {
				firstSlashIndex = i
			}
			break
		}
	}

	if atIndex == -1 || colonIndex == -1 || firstSlashIndex == -1 {
		return r, errors.New("invalid record")
	}

	r = Record{
		Account:     record[:atIndex],
		Driver:      record[atIndex+1 : colonIndex],
		Collection:  record[colonIndex+1 : firstSlashIndex],
		Record:      record[firstSlashIndex+1:],
		PoolPointer: record[:firstSlashIndex],
	}

	return r, nil
}
