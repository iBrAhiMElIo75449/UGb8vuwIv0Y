// 代码生成时间: 2025-09-02 05:39:59
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo/worker"
    "github.com/markbates/pkger"
)

// ThemeSwitcher is a worker that allows for theme switching.
type ThemeSwitcher struct {
    // CurrentTheme holds the current theme name.
    CurrentTheme string
}

// NewThemeSwitcher returns a new ThemeSwitcher worker.
func NewThemeSwitcher() *ThemeSwitcher {
    return &ThemeSwitcher{
        CurrentTheme: "default",
    }, nil
}

// SetTheme sets the current theme based on the provided theme name.
func (t *ThemeSwitcher) SetTheme(themeName string) error {
    // Check if the theme exists in available themes.
    if !t.isValidTheme(themeName) {
        return buffalo.NewError("Theme not found")
    }
    t.CurrentTheme = themeName
    return nil
}

// isValidTheme checks if the provided theme name is valid.
func (t *ThemeSwitcher) isValidTheme(themeName string) bool {
    // List of available themes.
    availableThemes := []string{"default", "dark", "light"}
    for _, theme := range availableThemes {
        if theme == themeName {
            return true
        }
    }
    return false
}

// Run is called by the Buffalo framework to run the worker.
func (t *ThemeSwitcher) Run(req *buffalo.Request) error {
    // Get the theme from the query parameters.
    themeName := req.Query("theme")
    if themeName == "" {
        themeName = "default"
    }
    // Set the theme.
    if err := t.SetTheme(themeName); err != nil {
        return err
    }
    // Store the theme in the session for persistence.
    if err := req.Session().Set("current_theme", themeName); err != nil {
        return err
    }
    req.Set("current_theme", t.CurrentTheme)
    return nil
}

func main() {
    app := buffalo.Automatic(buffalo.Options{
        AppName: "ThemeSwitcher",
        // Add ThemeSwitcher worker to the action pipeline.
        Middleware: []buffalo.MiddlewareFunc{
            buffalo.WrapHandlerFunc(worker.Worker(NewThemeSwitcher())),
        },
    })
    // The main application handler.
    app.GET("/", func(c buffalo.Context) error {
        c.Set("current_theme", c.Value("current_theme").(string))
        return c.Render(200, buffalo.HTML("index.html"))
    })
    // Start the Buffalo application.
    app.Serve()
}