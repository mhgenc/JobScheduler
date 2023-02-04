package main

import (
	"fmt"
	"time"
)

// Job represents a task to be executed in the future
type Job struct {
	ID       int
	Name     string
	Interval int
	Execute  func()
}

// JobQueue is a queue of Jobs
type JobQueue struct {
	Jobs []*Job
}

// NewJobQueue creates a new instance of JobQueue
func NewJobQueue() *JobQueue {
	return &JobQueue{
		Jobs: make([]*Job, 0),
	}
}

// Enqueue adds a new Job to the Queue
func (jq *JobQueue) Enqueue(job *Job) {
	jq.Jobs = append(jq.Jobs, job)
}

// Start starts executing the Jobs in the JobQueue
func (jq *JobQueue) Start() {
	for _, job := range jq.Jobs {
		go func(job *Job) {
			ticker := time.NewTicker(time.Duration(job.Interval) * time.Second)
			for range ticker.C {
				job.Execute()
			}
		}(job)
	}
}

func main() {
	jq := NewJobQueue()

	jq.Enqueue(&Job{
		ID:       1,
		Name:     "Task 1",
		Interval: 3,
		Execute: func() {
			fmt.Println("Executing Task 1 at", time.Now())
		},
	})

	jq.Enqueue(&Job{
		ID:       2,
		Name:     "Task 2",
		Interval: 5,
		Execute: func() {
			fmt.Println("Executing Task 2 at", time.Now())
		},
	})

	jq.Start()

	// Wait for all Jobs to complete
	select {}
}
