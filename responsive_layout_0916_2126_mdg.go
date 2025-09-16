// 代码生成时间: 2025-09-16 21:26:53
package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware"
    "github.com/gobuffalo/buffalo/worker"
    "github.com/gobuffalo/envy"
    "github.com/gobuffalo/packd"
    "github.com/gobuffalo/packr/v2"
    "log"
    "os"
)

// App is the Buffalo application.
var App *buffalo.App

// AppModels is where you define your Buffalo models.
var AppModels interface{}

// This is where you can define middleware for your application.
// Middleware is a function that Buffalo calls between routing and rendering.
// It's a great place to put Cross-Cutting Concerns.
func main() {
    env := envy.MustGet("GO_ENV", "development")
    if env == "development" {
        // In development, we turn on debugging and enable the live reload.
        app := buffalo.New(buffalo.Options{
           PrefetchTemplates: true,
        })
        app.Use(middleware.ParameterLogger)
        app.Use(middleware.Recover)
        app.ServeFiles("/assets", packr.New("myApp:assets", "assets"))
        App = app
    } else {
        // In production, we turn it off and set a custom logger.
        logLvl := envy.MustGet("LOG_LEVEL", "info")
        app := buffalo.New(buffalo.Options{
            Logger: &middleware.LogDispatcher{
                Level: logLvl,
                Out:   os.Stdout,
            },
        })
        app.Use(middleware.ParameterLogger)
        app.Use(middleware.Recover)
        app.ServeFiles("/assets", packr.New("myApp:assets", "assets"))
        App = app
    }

    // Here we带给App添加更多的中间件和路由器
    // ...

    // Finally, we call the Listen method to start the application listening for requests.
    if err := App.Start(); err != nil {
        log.Fatal(err)
    }
}

// NewHandler returns a new Buffalo handler.
func NewHandler(data worker.Worker, renderer buffalo.Renderer) buffalo.Handler {
    return func(c buffalo.Context) error {
        // Here you would add your logic to handle the request and
        // return an error if needed.
        // You can also use the data worker to interact with your database.
        // If your application needs to handle forms you can use the `buffalo.Populate` function
        // to automatically populate the form.
        // For example:
        // model := Model{}
        // if err := buffalo.Populate(c.Request(), &model); err != nil {
        //     return c.Error(500, err)
        // }
        // ...
        // Your logic here.
        //
        // If everything went well, you can render a template or a JSON response.
        // For example, rendering a template:
        // return renderer.Render(c, r, "template.html")
        // Or returning a JSON response:
        // return c.Render(200, json.Render(model))
        return nil
    }
}
