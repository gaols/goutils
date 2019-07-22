package goutils

import (
	"time"
)

// BeginningOfDate returns the beginning date of a time.
func BeginningOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// BeginningOfThisWeek returns the beginning date of this week.
func BeginningOfThisWeek() time.Time {
	return BeginningOfWeek(time.Now())
}

// BeginningOfWeek return the first date of the week where t belongs to.
func BeginningOfWeek(t time.Time) time.Time {
	wd := int(t.Weekday())
	if wd == 0 {
		wd = 7
	}
	d := t.AddDate(0, 0, -wd + 1)
	return BeginningOfDate(d)
}

// EndingOfDate returns the beginning date of a time.
func EndingOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// Yesterday returns the beginning date of yesterday.
func Yesterday() time.Time {
	today := Today()
	return today.AddDate(0, 0, -1)
}

// Today returns the beginning date of today.
func Today() time.Time {
	now := time.Now()
	return BeginningOfDate(now)
}

// Tomorrow returns the beginning date of tomorrow.
func Tomorrow() time.Time {
	today := Today()
	return today.AddDate(0, 0, 1)
}

// DaysAgo returns the time subtract specified days.
func DaysAgo(t time.Time, days int) time.Time {
	if days < 0 {
		panic("days should greater than or equals 0")
	}

	return t.AddDate(0, 0, -days)
}

// DaysAfter returns the time add specified days.
func DaysAfter(t time.Time, days int) time.Time {
	if days < 0 {
		panic("days should greater than or equals 0")
	}

	return t.AddDate(0, 0, days)
}

// DaysBetween calc days between two days.
// goutils.DaysBetween(goutils.Yesterday(), goutils.Today()) = 1
// goutils.DaysBetween(goutils.Today(), goutils.Yesterday()) = -1
func DaysBetween(startTime, endTime time.Time) int {
	diff := BeginningOfDate(endTime).Sub(BeginningOfDate(startTime))
	return int(diff.Hours() / 24)
}
