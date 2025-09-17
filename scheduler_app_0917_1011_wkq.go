// 代码生成时间: 2025-09-17 10:11:19
package main

import (
    "time"
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware"
    "github.com/gobuffalo/buffalo/worker"
    "github.com/robfig/cron/v3"
)

// SchedulerApp is the main application struct
type SchedulerApp struct {
    *buffalo.App
    Cron *cron.Cron
}

// NewSchedulerApp creates a new instance of SchedulerApp
func NewSchedulerApp() *SchedulerApp {
    app := buffalo.New(buffalo.Options{})

    // Add middleware
    app.Use(middleware.Logger)
    app.Use(middleware.Recovery)

    // Initialize cron scheduler
    cron := cron.New()
    app.Injector().RegisterSingleton(cron)

    return &SchedulerApp{
        App: app,
        Cron: cron,
    }
}

// Add a job to the scheduler
func (app *SchedulerApp) AddJob(spec string, cmd worker.Command) {
    _, err := app.Cron.AddFunc(spec, func() { _ = app.Worker(cmd) })
    if err != nil {
        buffalo.AppLog(app).Error(err)
    }
}

// Schedule starts the scheduler and runs the Buffalo application
func (app *SchedulerApp) Schedule() {
    app.Cron.Start()
    defer app.Cron.Stop()
    app.ServeFiles("/public", "./public")
}

// Main is the entry point of the application
func main() {
    app := NewSchedulerApp()
    defer app.Close()

    // Example of adding a job to run every minute
    app.AddJob("* * * * *", worker.CommandFunc(func() error {
        // Place your job logic here
        buffalo.AppLog(app).Info("Job executed at ", time.Now())
        return nil
    }))

    // Start the application
    app.Schedule()
}
