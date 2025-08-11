// 代码生成时间: 2025-08-12 00:32:03
This program demonstrates how to create a scheduler that can execute tasks at specific times or intervals.
It follows Go's best practices and is designed to be maintainable and extensible.
*/

package main

import (
    "log"
    "time"

    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/worker"
    "github.com/robfig/cron/v3"
)

// Scheduler defines the application's scheduler
type Scheduler struct {
    c *cron.Cron
}

// NewScheduler creates a new instance of Scheduler
func NewScheduler() *Scheduler {
    return &Scheduler{
        c: cron.New(cron.WithSeconds()),
    },
}

// AddJob adds a new job to the scheduler
func (s *Scheduler) AddJob(spec string, cmd func()) error {
    _, err := s.c.AddFunc(spec, cmd)
    return err
}

// Start starts the scheduler
func (s *Scheduler) Start() {
    s.c.Start()
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
    s.c.Stop()
}

// JobHandler is the handler that will be called by the scheduler
func JobHandler(c buffalo.Context) error {
    // Your job logic here
    log.Println("JobHandler executed")
    return nil
}

func main() {
    app := buffalo.Automatic()

    // Create a new scheduler
    sch := NewScheduler()

    // Add a job to run every minute
    if err := sch.AddJob("* * * * * *", JobHandler); err != nil {
        log.Fatalf("Failed to add job: %s", err)
    }

    // Start the scheduler in the background
    go sch.Start()

    // Run the app
    if err := app.Serve(); err != nil {
        log.Fatal(err)
    }

    // Stop the scheduler when the app is shutting down
    sch.Stop()
}
