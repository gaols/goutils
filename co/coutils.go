package co

import "sync"

// Job abstract a job which can be executed by PoolJobs.
type Job interface {
	// Run run the current job
	Run() (ret interface{}, err error)
}

// PoolJobs execute jobs concurrently with limited concurrency, you can stop on error or ignore errors and proceed on.
func PoolJobs(jobs []Job, concurrency int, stopOnErr bool, retHandler func(job Job, ret interface{}, err error)) {
	if jobs == nil || len(jobs) <= 0 {
		return
	}

	mux := &sync.Mutex{}
	hasErr := false

	sem := make(chan bool, concurrency)
	for _, job := range jobs {
		mux.Lock()
		if hasErr && stopOnErr {
			mux.Unlock()
			break
		}
		mux.Unlock()

		sem <- true
		go func(j Job) {
			defer func() { <-sem }()
			ret, err := j.Run()
			if err != nil {
				mux.Lock()
				hasErr = true
				mux.Unlock()
			}
			retHandler(j, ret, err)
		}(job)
	}

	for i := 0; i < concurrency; i++ {
		sem <- true
	}
}
