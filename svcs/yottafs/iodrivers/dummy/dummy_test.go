package dummy

import (
	"testing"
	"yottafs/iodrivers"
)

func TestDummy(t *testing.T) {

	d, err := New()
	if err != nil {
		t.Fatal("Failed instantiating driver", err)
	}

	req := iodrivers.Request{
		Record: "Test",
	}

	_, err = d.Create(req)
	if err != nil {
		t.Fatal("Error creating file", err)
	}

	_, err = d.Read(req)
	if err != nil {
		t.Fatal("Error reading file", err)
	}

	_, err = d.Update(req)
	if err != nil {
		t.Fatal("Error updating file", err)
	}

	err = d.Delete(req)
	if err != nil {
		t.Fatal("Error deleting file", err)
	}

}
