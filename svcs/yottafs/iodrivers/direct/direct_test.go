package direct

import (
	"fmt"
	"golang.org/x/sys/unix"
	"testing"
)

func TestDirect(t *testing.T) {

	d, err := New("/tmp/testDirect")
	if err != nil {
		t.Fatal("Failed instantiating driver", err)
	}

	fmt.Println(d)

	_, err = New("/tmp/testDirect")
	if err != nil {
		t.Fatal("Failed instantiating driver a second time", err)
	}

	_, err = New("/")
	if err != unix.EACCES {
		t.Fatal("Permissions should have failed", err)
	}

}
