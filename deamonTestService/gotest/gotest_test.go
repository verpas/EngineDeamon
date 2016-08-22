package gotest

import "testing"

func TestFunc(t *testing.T) {
	if Print(1) {
		t.Fail()
	}
}
