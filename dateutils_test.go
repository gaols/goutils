package goutils

import "testing"

func TestYesterday(t *testing.T) {
	now := Today()
	yesterday := now.AddDate(0, 0, -1)
	if !Yesterday().Equal(yesterday) {
		t.FailNow()
	}
}

func TestTomorrow(t *testing.T) {
	now := Today()
	tomorrow := now.AddDate(0, 0, 1)
	if !Tomorrow().Equal(tomorrow) {
		t.FailNow()
	}
}
