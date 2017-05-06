package pfx

import (
	"fmt"
	"testing"
)

func TestE(t *testing.T) {
	expect := "pfx.eTest: eTest message"
	if err := eTest(); err.Error() != expect {
		t.Error("Got", err.Error(), "expected", expect)
	}
}

func TestENil(t *testing.T) {
	expect := error(nil)
	if err := nilETest(); err != expect {
		t.Error("Got", err.Error(), "expected nil")
	}
}

func TestES(t *testing.T) {
	expect := "pfx.esTest: esTest message"
	if err := esTest(); err.Error() != expect {
		t.Error("Got", err.Error(), "expected", expect)
	}
}

func TestDirectES(t *testing.T) {
	expect := "pfx.TestDirectES: Nope"
	if err := Err("Nope"); err.Error() != expect {
		t.Error("Got", err.Error(), "expected", expect)
	}
}

func TestFullPathE(t *testing.T) {
	defer func() {
		FullyQualifiedPath = false
	}()

	FullyQualifiedPath = true

	expect := "github.com/carbocation/pfx.esTest: esTest message"
	if err := esTest(); err.Error() != expect {
		t.Error("Got", err.Error(), "expected", expect)
	}
}

func eTest() error {
	return Err(fmt.Errorf("eTest message"))
}

func esTest() error {
	return Err("esTest message")
}

func nilETest() error {
	return Err(nil)
}
