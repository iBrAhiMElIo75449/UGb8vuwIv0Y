// 代码生成时间: 2025-08-24 20:36:41
package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/generators/assets/css"
    "github.com/gobuffalo/buffalo/generators/assets/js"
    "github.com/gobuffalo/buffalo/generators/apps/action"
    "github.com/gobuffalo/buffalo/meta/tags"
    "github.com/gobuffalo/envy"
    "log"
    "net/http"
)

// 应用的主结构
type App struct {
    *buffalo.App
}

// New 用于初始化 Buffalo 应用
func New() *App {
    if env := envy.Get("BUFFALO_ENV", "development"); env == "development" {
        buffalo.Debug = true
    }
    app := buffalo.New(buffalo.Options{
        PrettyErrors: true,
    })

    // 定义路由
    app.GET("/", HomeHandler)
    app.GET("/about", AboutHandler)

    // 响应式布局的 CSS 和 JS 文件
    app.Use("/assets/*path", css.Styles)
    app.Use("/assets/*path", js.Scripts)

    return &App{App: app}
}

// HomeHandler 处理根路径的请求
func HomeHandler(c buffalo.Context) error {
    return c.Render(200, r.HTML("home.html"))
}

// AboutHandler 处理关于页面的请求
func AboutHandler(c buffalo.Context) error {
    return c.Render(200, r.HTML("about.html"))
}

// main 函数启动 Buffalo 应用
func main() {
    if err := New().Serve(); err != nil {
        log.Fatal(err)
    }
}
