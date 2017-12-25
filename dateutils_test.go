package goutils

import (
	"testing"
	"time"
)

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

func TestDaysAgo(t *testing.T) {
	now := time.Now()
	if !BeginningOfDate(DaysAgo(now, 1)).Equal(Yesterday()) {
		t.FailNow()
	}
}

func TestDaysAfter(t *testing.T) {
	now := time.Now()
	if !BeginningOfDate(DaysAfter(now, 1)).Equal(Tomorrow()) {
		t.FailNow()
	}
}

func TestEndingOfDate(t *testing.T) {
	now := time.Now()
	end := EndingOfDate(now)
	if end.Hour() != 23 || end.Minute() != 59 || end.Second() != 59 || end.Year() != now.Year() || end.Month() != now.Month() || end.Day() != now.Day() {
		t.FailNow()
	}
}
