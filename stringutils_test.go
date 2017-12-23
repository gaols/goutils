package goutils

import "testing"

func TestIsBlank(t *testing.T) {
	if !IsBlank("") {
		t.FailNow()
	}
	if !IsBlank(" ") {
		t.FailNow()
	}
	if !IsBlank(" \t\n") {
		t.FailNow()
	}
	if IsBlank("h") {
		t.FailNow()
	}
}
