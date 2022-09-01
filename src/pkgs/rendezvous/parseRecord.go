package rendezvous

import (
	"errors"
)

type ParsedRecord struct {
	Account          string
	Driver           string
	TableName        string
	RecordName       string
	TableIdentifier  string
	RecordIdentifier string
	RecordRow        string
}

var (
	MalformedRecordErr = errors.New("Malformed Record")
)

func ParseRecord(recordString string) (ParsedRecord, error) {

	var parsedRecord ParsedRecord

	atIndex := -1
	columnCounter := 0
	columnIndexs := [2]int{0, len(recordString)}
	slashIndexs := make([]int, 0, 1)

	for idx, char := range recordString {

		switch char {
		case '@':
			if atIndex != -1 {
				return parsedRecord, MalformedRecordErr
			}
			atIndex = idx
			columnIndexs[0] = idx
		case ':':
			if columnCounter > 1 {
				return parsedRecord, MalformedRecordErr
			}
			columnIndexs[columnCounter] = idx
			columnCounter++
		case '/':
			slashIndexs = append(slashIndexs, idx)
		}

	}

	if atIndex == -1 {
		return parsedRecord, MalformedRecordErr
	}

	if len(slashIndexs) == 0 {
		return parsedRecord, MalformedRecordErr

	}

	Account := recordString[:atIndex]
	Driver := recordString[atIndex+1 : columnIndexs[0]]
	TableName := recordString[columnIndexs[0]+1 : slashIndexs[0]]
	RecordName := recordString[slashIndexs[0]+1 : columnIndexs[1]]
	TableIdentifier := recordString[:slashIndexs[0]]
	RecordIdentifier := recordString[:columnIndexs[1]]
	RecordRow := recordString[columnIndexs[1]+1 : len(recordString)]

	if Driver == "" {
		Driver = "kv"
	}

	parsedRecord = ParsedRecord{
		Account,
		Driver,
		TableName,
		RecordName,
		TableIdentifier,
		RecordIdentifier,
		RecordRow,
	}

	return parsedRecord, nil
}
