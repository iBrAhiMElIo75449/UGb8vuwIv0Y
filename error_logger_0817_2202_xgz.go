// 代码生成时间: 2025-08-17 22:02:03
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "github.com/gobuffalo/buffalo/generators/claims"    // 用于日志生成
    "log"
    "os"
)

// ErrorLogger 定义错误日志收集器的结构
type ErrorLogger struct {
    Logger *log.Logger
}

// NewErrorLogger 创建并返回一个新的 ErrorLogger 实例
func NewErrorLogger() *ErrorLogger {
    // 创建日志文件
    file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    // 创建 logger 实例
    logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
    return &ErrorLogger{Logger: logger}
}

// LogError 记录错误信息到日志文件
func (e *ErrorLogger) LogError(err error) {
    if err != nil {
        e.Logger.Println(err)
    }
}

// App 定义 Buffalo 应用结构
type App struct {
    *buffalo.App
    Logger *ErrorLogger
}

// NewApp 创建并返回一个新的 App 实例
func NewApp() *App {
    a := buffalo.NewApp(
        buffalo.Options{
            AppName: "error-logger", // 应用名称
        },
    )
    a.Use(
        generators.NewLogger, // 使用 Buffalo 的内置日志器
    )
    e := NewErrorLogger()
    a.Use(func(next buffalo.Handler) buffalo.Handler {
        return func(c buffalo.Context) error {
            err := next(c)
            if err != nil {
                e.LogError(err) // 记录错误到自定义日志器
            }
            return err
        }
    })
    return &App{App: a, Logger: e}
}

// main 函数是程序的入口点
func main() {
    app := NewApp()
    if err := app.Serve(); err != nil {
        app.Logger.LogError(err) // 如果启动服务失败，记录错误信息
    }
}