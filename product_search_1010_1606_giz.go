// 代码生成时间: 2025-10-10 16:06:00
package main

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/pop"
    "github.com/gobuffalo/buffalo/middleware"
    "net/http"
    "log"
    "encoding/json"
)

// Product 代表一个商品
type Product struct {
    ID        uint   "json:-"
    Name      string "json:"name"
    Price     float64 "json:"price"
    CreatedAt string  "json:"createdAt"
    UpdatedAt string  "json:"updatedAt"
}

// ProductSearchHandler 处理搜索请求
func ProductSearchHandler(c buffalo.Context) error {
    // 获取查询参数
    query := c.Param("query")
    if query == "" {
        // 如果没有查询参数，返回错误
        return buffalo.NewError("Query parameter 'query' is required", http.StatusBadRequest)
    }

    // 搜索商品
    products, err := searchProducts(query)
    if err != nil {
        // 如果搜索失败，返回错误
        return err
    }

    // 返回搜索结果
    return c.Render(http.StatusOK, r.JSON(products))
}

// searchProducts 搜索商品
func searchProducts(query string) ([]Product, error) {
    // 初始化数据库连接
    tx := buffaloApp.DB.Begin()
    defer tx.Rollback()

    // 执行搜索查询
    var products []Product
    if err := tx.Where("name ILIKE ?