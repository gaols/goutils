package co

import (
	"testing"
	"errors"
)

type A struct {
	Data int
}

func (a *A) Run() (interface{}, error) {
	if a.Data%2 == 0 {
		return "", errors.New("test error")
	}
	return "abc", nil
}

func TestPoolJobs(t *testing.T) {
	jobs := []Job{&A{Data: 2}, &A{Data: 1}, &A{Data: 0}, &A{Data: 4}, &A{Data: 6}, &A{Data: 7}, &A{Data: 8}, &A{Data: 10}}

	PoolJobs(jobs, 5, false, func(job Job, ret interface{}, err error) {
		t.Log("job", job, "done,ret=", ret, ",error=", err)
	})
}
