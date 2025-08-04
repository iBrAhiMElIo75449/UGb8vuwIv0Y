// 代码生成时间: 2025-08-05 02:27:46
package main

import (
	"os"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/generators"
	"github.com/gobuffalo/buffalo/generators/app"
	"github.com/gobuffalo/buffalo/generators/assets/css/sass"
	"github.com/gobuffalo/buffalo/generators/assets/js/npm"
)

// main is the entry point for the Buffalo application.
func main() {
	// Initialize a new Buffalo application.
	app := buffalo.Automatic(buffalo.Options{
		AppName:       "ResponsiveLayout",
		ForceOverwrite: true,
	})

	// Generate the necessary files for CSS and JS.
	if err := app.Generators(app.Options).Add(app.Generators().New(sass.Generator{})).Run(); err != nil {
		app.Stop(err)
	}
	if err := app.Generators(app.Options).Add(app.Generators().New(npm.Generator{})).Run(); err != nil {
		app.Start()
	}

	// Start the Buffalo application.
	if err := app.Serve(); err != nil && err != http.ErrServerClosed {
		app.Stop(err)
	} else {
		app.Stop(buffalo.ExitNormal)
	}
}
