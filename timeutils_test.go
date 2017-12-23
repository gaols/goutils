package goutils

import (
	"testing"
	"time"
)

func TestFmtTime(t *testing.T) {
	ti, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-12-23 12:06:06", time.Now().Location())
	ft, err := FmtTime(ti, "-datetime")
	if err != nil {
		t.FailNow()
	}
	if ft != "2017-12-23 12:06:06" {
		t.FailNow()
	}
	ft, err = FmtTime(ti, "-datetime-")
	if err != nil {
		t.FailNow()
	}
	if ft != "2017-12-23 12:06" {
		t.FailNow()
	}
	ft, err = FmtTime(ti, "-datetime--")
	if err != nil {
		t.FailNow()
	}
	if ft != "2017-12-23 12" {
		t.FailNow()
	}
	ft, err = FmtTime(ti, "-date")
	if err != nil {
		t.FailNow()
	}
	if ft != "2017-12-23" {
		t.FailNow()
	}
	ft, err = FmtTime(ti, "time")
	if err != nil {
		t.FailNow()
	}
	if ft != "12:06:06" {
		t.FailNow()
	}

}

func TestMustFmtTime(t *testing.T) {
	ti, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-12-23 12:06:06", time.Now().Location())
	ft := MustFmtTime(ti, "/datetime")
	if ft != "2017/12/23 12:06:06" {
		t.FailNow()
	}
}
