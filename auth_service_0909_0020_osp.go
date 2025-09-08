// 代码生成时间: 2025-09-09 00:20:48
package main

import (
    "buffalo"
    "github.com/gobuffalo/buffalo/worker"
    "github.com/markbates/going/defaults"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// AuthService 结构体，用于用户身份认证
type AuthService struct {
    DB *gorm.DB
}

// NewAuthService 创建 AuthService 实例
func NewAuthService(db *gorm.DB) *AuthService {
    return &AuthService{DB: db}
}

// Authenticate 用户身份认证方法
func (as *AuthService) Authenticate(c buffalo.Context) error {
    // 从上下文获取用户名和密码
    username := c.Value("username")
    password := c.Value("password")

    // 检查用户名和密码是否为空
    if username == nil || password == nil {
        return buffalo.NewError("Username or password is required")
    }

    // 从数据库查找用户
    user := User{}
    if err := as.DB.Where(Username+" = ? AND password = ?", username, password).First(&user).Error; err != nil {
        // 如果用户不存在或密码错误，返回错误
        return buffalo.NewError("Invalid username or password")
    }

    // 如果用户存在并且密码正确，设置认证令牌
    token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })
    if err != nil {
        return err
    }
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return err
    }

    // 将认证令牌设置到上下文中
    c.Set("Authorization", "Bearer " + tokenString)
    return nil
}

// User 定义用户模型
type User struct {
    gorm.Model
    Username string `gorm:"unique" json:"username"`
    Password string `json:"password"`
}

// main 函数，初始化应用和数据库
func main() {
    app := buffalo.Automatic()

    // 设置数据库连接
    db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // 自动迁移数据库模式
    db.AutoMigrate(&User{})

    // 创建 AuthService 实例
    authService := NewAuthService(db)

    // 定义认证路由
    app.GET("/login", func(c buffalo.Context) error {
        // 从请求中解析用户名和密码
        username := c.Request().FormValue("username")
        password := c.Request().FormValue("password")
        c.Set("username", username)
        c.Set("password", password)

        // 调用认证方法
        if err := authService.Authenticate(c); err != nil {
            return c.Error(401, err)
        }

        // 返回成功的响应
        return c.Render(200, r.JSON(map[string]string{
            "message": "Login successful"
        }))
    })

    // 启动应用
    app.Serve()
}
