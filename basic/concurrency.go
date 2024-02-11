package basic

import (
	"fmt"
	"sync"
	"time"
)

func ConcurrencyExample() {
	var jobList [10]Job

	for i := 0; i < len(jobList); i++ {
		jobList[i] = &SquireJob{index: i}

	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			defer wg.Done()
			job.Do()
		}()
	}
	wg.Wait()
}

type Job interface {
	Do()
}

type SquireJob struct {
	index int
}

func (j *SquireJob) Do() {
	fmt.Printf("%d worker start", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d worker end", j.index)
}
