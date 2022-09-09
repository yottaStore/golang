package methods

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestMethods(t *testing.T) {

	writePayload := []byte("hello wordl")
	writePath := "/tmp/yottatest/test.txt"

	_, err := Write("/tmp/kwa/test.txt", writePayload, false)
	if err == nil {
		t.Error("Write should have failed", err)
	}

	writeResp, err := Write(writePath, writePayload, true)
	if err != nil {
		t.Fatal("Error while writing", err)
	}

	readResp, err := Read(writePath)
	if err != nil {
		t.Fatal("Error while reading", err)
	}

	str1 := string(writePayload)
	str2 := string(readResp.Data[:len(writePayload)])

	if str1 != str2 {
		t.Error("Data not matching")
	}
	if writeResp.AbaToken != readResp.AbaToken {
		t.Error("Aba counter not matching")
	}

	err = Delete(writePath)
	if err != nil {
		t.Error("Error deleting", err)
	}

	_, err = Read(writePath)
	if err != unix.ENOENT {
		t.Fatal("File shouldn't exist", err)
	}

}
