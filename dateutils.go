package goutils

import "time"

// BeginningOfDate returns the beginning date of a time.
func BeginningOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// Yesterday returns the beginning date of yesterday.
func Yesterday() time.Time {
	now := Today()
	return now.AddDate(0, 0, -1)
}

// Today returns the beginning date of today.
func Today() time.Time {
	now := time.Now()
	return BeginningOfDate(now)
}

// Tomorrow returns the beginning date of tomorrow.
func Tomorrow() time.Time {
	now := Today()
	return now.AddDate(0, 0, 1)
}
