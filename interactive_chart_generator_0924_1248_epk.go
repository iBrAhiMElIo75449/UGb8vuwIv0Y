// 代码生成时间: 2025-09-24 12:48:13
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/generators/assets/generators/action"
    "github.com/gobuffalo/buffalo/generators/assets/generators/resource"
    "github.com/markbates/inflect"
    "strings"
)

// InteractiveChartGenerator represents the generator for creating an interactive chart
type InteractiveChartGenerator struct {
    *action.Generator
}

// New creates a new instance of InteractiveChartGenerator
func New(opts *buffalo.GeneratorOptions) (*InteractiveChartGenerator, error) {
    g := &InteractiveChartGenerator{
        Generator: &action.Generator{
            GOPkg: opts.AppName,
        },
    }
    return g, nil
}

// Run creates the necessary files for an interactive chart
func (g *InteractiveChartGenerator) Run() error {
    // Define the resource name and path
    resourceName := strings.Title(g.Name)
    resourcePath := inflect.Pluralize(resourceName)

    // Generate the model file
    if err := g.generateModelFile(resourceName, resourcePath); err != nil {
        return err
    }

    // Generate the controller file
    if err := g.generateControllerFile(resourceName, resourcePath); err != nil {
        return err
    }

    // Generate the view files
    if err := g.generateViewFiles(resourceName, resourcePath); err != nil {
        return err
    }

    // Generate the route file
    if err := g.generateRouteFile(resourceName, resourcePath); err != nil {
        return err
    }

    // Generate the test files
    if err := g.generateTestFiles(resourceName, resourcePath); err != nil {
        return err
    }

    return nil
}

// generateModelFile creates the model file for the resource
func (g *InteractiveChartGenerator) generateModelFile(resourceName, resourcePath string) error {
    // Create a new model generator
    modelGen := resource.NewModel(g.GOPkg, g.Name, g.Name)

    // Generate the model file
    if err := modelGen.Run(); err != nil {
        return err
    }

    return nil
}

// generateControllerFile creates the controller file for the resource
func (g *InteractiveChartGenerator) generateControllerFile(resourceName, resourcePath string) error {
    // Create a new controller generator
    ctrlGen := resource.NewController(g.GOPkg, resourcePath, g.Name)

    // Generate the controller file
    if err := ctrlGen.Run(); err != nil {
        return err
    }

    return nil
}

// generateViewFiles creates the view files for the resource
func (g *InteractiveChartGenerator) generateViewFiles(resourceName, resourcePath string) error {
    // Create a new view generator
    viewGen := resource.NewView(g.GOPkg, resourcePath, g.Name)

    // Generate the view files
    if err := viewGen.Run(); err != nil {
        return err
    }

    return nil
}

// generateRouteFile creates the route file for the resource
func (g *InteractiveChartGenerator) generateRouteFile(resourceName, resourcePath string) error {
    // Define the route file path
    routeFilePath := "./actions/app/" + resourcePath + ".go"

    // Create a new route generator
    routeGen := resource.NewRoute(g.GOPkg, resourcePath, g.Name)

    // Generate the route file
    if err := routeGen.Run(); err != nil {
        return err
    }

    return nil
}

// generateTestFiles creates the test files for the resource
func (g *InteractiveChartGenerator) generateTestFiles(resourceName, resourcePath string) error {
    // Create a new test generator
    testGen := resource.NewTest(g.GOPkg, resourcePath, g.Name)

    // Generate the test files
    if err := testGen.Run(); err != nil {
        return err
    }

    return nil
}

func main() {
    // Create a new generator options
    opts := &buffalo.GeneratorOptions{
        AppName: "InteractiveChartApp",
        Name: "chart",
    }

    // Create a new instance of InteractiveChartGenerator
    g, err := New(opts)
    if err != nil {
        panic(err)
    }

    // Run the generator
    if err := g.Run(); err != nil {
        panic(err)
    }
}
