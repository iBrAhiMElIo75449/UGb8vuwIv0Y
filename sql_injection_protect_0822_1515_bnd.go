// 代码生成时间: 2025-08-22 15:15:43
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"
    "github.com/markbates/buffalo"
    "github.com/markbates/buffalo/x/buffaloctl/cmd"
    "github.com/markbates/pop/v5"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// main 是程序的入口点
func main() {
    // 初始化buffalo应用
    app := buffalo.New(buffalo.Options{})

    // 设置数据库连接
    app.ServeFiles("/assets", assetsPath())
    db := setupDB(app)
    defer db.Close()

    // 定义路由
    app.GET("/", homeHandler(db))
    app.Serve()
}

// User 模型定义
type User struct {
    gorm.Model
    Name string
}

// HomeHandler 处理主页请求，演示防止SQL注入
func homeHandler(db *sql.DB) buffalo.HandlerFunc {
    return func(c buffalo.Context) error {
        // 从请求中获取参数
        name := c.Param("name")

        // 构造查询，使用参数化查询防止SQL注入
        var user User
        result := db.QueryRow(`SELECT * FROM users WHERE name = ?`, name).Scan(&user)
        if result.Err() != nil {
            return result.Err()
        }

        // 返回结果
        return c.Render(200, r.JSON(map[string]string{"message": "User found"}))
    }
}

// setupDB 设置数据库连接
func setupDB(app *buffalo.App) *sql.DB {
    // 配置数据库连接字符串
    dsn := fmt.Sprintf("%s?mode=rwc&cache=shared&_fk=1", os.Getenv("DATABASE_URL"))

    // 打开数据库连接
    db, err := sql.Open("sqlite3", dsn)
    if err != nil {
        log.Fatal(err)
    }

    // 应用数据库迁移
    if err := migrateDatabase(db); err != nil {
        log.Fatal(err)
    }

    return db
}

// migrateDatabase 应用数据库迁移
func migrateDatabase(db *sql.DB) error {
    dialect := pop.NewSqliteDialect(DB)
    if err := dialect.Connection().Ping(); err != nil {
        return err
    }

    // 自动迁移模型
    err := dialect.AutoMigrate(&User{})
    if err != nil {
        return err
    }

    return nil
}
