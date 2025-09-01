// 代码生成时间: 2025-09-01 10:13:26
package main

import (
    "buffalo"
    "buffalo/worker"
    "log"
    "time"
)

// Scheduler represents a job scheduler
type Scheduler struct {
    interval time.Duration
    jobs     []func() error
}

// NewScheduler creates a new scheduler with an interval
func NewScheduler(interval time.Duration) *Scheduler {
    return &Scheduler{
        interval: interval,
    }
}

// AddJob adds a new job to the scheduler
func (s *Scheduler) AddJob(job func() error) {
    s.jobs = append(s.jobs, job)
}

// Run starts the scheduler and runs jobs at the specified interval
func (s *Scheduler) Run() {
    for {
        for _, job := range s.jobs {
            if err := job(); err != nil {
                log.Printf("Error running job: %v", err)
            }
        }
        time.Sleep(s.interval)
    }
}

// StartWorker starts a new worker for the scheduler
func StartWorker(interval time.Duration) {
    app := buffalo.NewApp(buffalo.Options{})
    defer app.Close()

    // Create a new scheduler with the given interval
    scheduler := NewScheduler(interval)

    // Add jobs to the scheduler
    scheduler.AddJob(func() error {
        // Example job: print a message
        log.Println("Running example job...")
        return nil
    })

    // Start the scheduler in a goroutine
    go scheduler.Run()

    // Start the Buffalo worker
    worker := worker.New(app)
    if err := worker.Start(); err != nil {
        log.Fatal(err)
    }
}

func main() {
    // Set the interval for the scheduler (e.g., 1 minute)
    interval := 1 * time.Minute

    // Start the worker with the scheduler
    StartWorker(interval)
}
