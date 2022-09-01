package main

import (
	"bytes"
	"testing"
	"yottaStore/yottaStore-go/src/libs/drivers/direct"
)

func TestDirectDriver(t *testing.T) {

	driver := direct.New()

	payload := []byte("Testing direct driver\n")
	if err := driver.Write("./test.txt", payload); err != nil {
		t.Error(err)
	}

	buff, err := driver.Read("./test.txt")
	if err != nil {
		t.Error(err)
	}
	if len(buff) != 4096 {
		t.Error("Buffer size wrong")
	}

	if res := bytes.Compare(payload, buff[:len(payload)]); res != 0 {
		t.Error("Written payload doesn't match")
	}

	appendPayload := []byte("Testing direct driver append\n")
	if err = driver.Append("./test.txt", appendPayload); err != nil {
		t.Error(err)
	}

	buff, err = driver.Read("./test.txt")
	if err != nil {
		t.Error(err)
	}
	if len(buff) != 4096 {
		t.Error("Buffer size wrong")
	}

	payload = append(payload, appendPayload...)
	if res := bytes.Compare(payload, buff[:len(payload)]); res != 0 {
		t.Error("Appended payload doesn't match")
	}

	err = driver.Delete("./test.txt")
	if err != nil {
		t.Error(err)
	}

	err = driver.Delete("./test.txt")
	if err == nil {
		t.Error("File not deleted!")
	}

}
