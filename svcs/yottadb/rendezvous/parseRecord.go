package rendezvous

import (
	"errors"
)

type ParsedRecord struct {
	Account              string
	Driver               string
	CollectionName       string
	RecordName           string
	CollectionIdentifier string
	CollectionPointer    string
	RecordIdentifier     string
	RecordRows           string
}

var (
	MalformedRecordErr = errors.New("Malformed Record")
)

// Example record
//
// account@[driver:]tableName/recordName[/recordRow[/subRecordRow]]
// account@tableName/recordName

func ParseRecord(recordString string) (ParsedRecord, error) {

	var parsedRecord ParsedRecord

	atIndex := -1
	colonIndex := -1
	slashIndexs := make([]int, 0, 1)

	for idx, char := range recordString {
		switch char {
		case '@':
			if atIndex != -1 {
				return parsedRecord, MalformedRecordErr
			}
			atIndex = idx
		case ':':
			if colonIndex != -1 {
				return parsedRecord, MalformedRecordErr
			}
			colonIndex = idx
		case '/':
			slashIndexs = append(slashIndexs, idx)
			if len(slashIndexs) == 2 {
				continue
			}
		}
	}

	if atIndex == -1 || len(slashIndexs) == 0 {
		return parsedRecord, MalformedRecordErr
	}

	// Should we return error on missing driver?
	if colonIndex == -1 {
		colonIndex = atIndex
		recordString = recordString[:atIndex+1] + "kv:" + recordString[atIndex+1:]
		colonIndex = atIndex + 3
		for idx, v := range slashIndexs {
			slashIndexs[idx] = v + 3
		}
	}

	if len(slashIndexs) == 1 {
		recordString = recordString + "/"
		slashIndexs = append(slashIndexs, len(recordString)-1)
	}

	Account := recordString[:atIndex]
	Driver := recordString[atIndex+1 : colonIndex]
	CollectionName := recordString[colonIndex+1 : slashIndexs[0]]
	RecordName := recordString[slashIndexs[0]+1 : slashIndexs[1]]
	CollectionIdentifier := "/" + Account + "/" + CollectionName
	CollectionPointer := "/" + Account + "-" + CollectionName
	RecordIdentifier := "/" + Account + "/" + recordString[colonIndex+1:slashIndexs[1]]
	RecordRows := recordString[slashIndexs[1]+1:]

	parsedRecord = ParsedRecord{
		Account,
		Driver,
		CollectionName,
		RecordName,
		CollectionIdentifier,
		CollectionPointer,
		RecordIdentifier,
		RecordRows,
	}

	return parsedRecord, nil
}
