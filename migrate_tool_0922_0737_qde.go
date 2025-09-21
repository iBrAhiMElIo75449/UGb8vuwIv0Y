// 代码生成时间: 2025-09-22 07:37:35
package main

import (
    "log"
    "os"
    "path/filepath"

    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/meta/migration"
    "github.com/gobuffalo/buffalo"
)

// main is the entry point for the program
func main() {
    // Parse command line arguments for migration command
    if len(os.Args) < 3 || os.Args[1] != "migrate" {
        log.Fatalf("Usage: %s migrate <command> <args>", filepath.Base(os.Args[0]))
    }

    // Create a new buffalo app
    app := buffalo.App()

    // Create a new migration generator
    mg, err := migration.New(migration.Options{
        App: app,
    })
    if err != nil {
        log.Fatalf("Error creating migration generator: %s", err)
    }

    // Check the command and perform the corresponding action
    switch os.Args[2] {
    case "up":
        err = mg.Up()
        if err != nil {
            log.Fatalf("Error running migration up: %s", err)
        }
    case "down":
        err = mg.Down()
        if err != nil {
            log.Fatalf("Error running migration down: %s", err)
        }
    case "create":
        if len(os.Args) < 4 {
            log.Fatalf("Usage: %s migrate create <name>", filepath.Base(os.Args[0]))
        }
        err = mg.Create(os.Args[3])
        if err != nil {
            log.Fatalf("Error creating migration: %s", err)
        }
    default:
        log.Fatalf("Unknown migration command: %s", os.Args[2])
    }
}
