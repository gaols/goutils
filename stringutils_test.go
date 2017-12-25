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

func TestIsEmpty(t *testing.T) {
	if !IsEmpty("") {
		t.FailNow()
	}
	if IsEmpty("a") {
		t.FailNow()
	}
}

func TestDefaultIfEmpty(t *testing.T) {
	if DefaultIfEmpty("", "foo") != "foo" {
		t.FailNow()
	}
}

func TestDefaultIfBlank(t *testing.T) {
	if DefaultIfBlank(" \t\r", "foo") != "foo" {
		t.FailNow()
	}
}
