// 代码生成时间: 2025-08-04 05:16:16
package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
)

// Renamer 定义批量重命名的工具
type Renamer struct {
    // 目录路径
    dir string
    // 重命名规则
    rule func(string) string
}

// NewRenamer 初始化 Renamer 实例
func NewRenamer(dir string, rule func(string) string) *Renamer {
    return &Renamer{
        dir: dir,
        rule: rule,
    }
}

// Rename 执行批量重命名操作
func (r *Renamer) Rename() error {
    err := filepath.WalkDir(r.dir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil
        }

        // 避免重命名目录
        if filepath.Ext(path) == "" {
            return nil
        }

        // 应用重命名规则
        newPath := r.rule(path)
        if newPath == path {
            return nil
        }

        // 重命名文件
        if err := os.Rename(path, newPath); err != nil {
            return err
        }
        fmt.Printf("Renamed '%s' to '%s'
", path, newPath)
        return nil
    })
    return err
}

// ExampleRule 演示重命名规则，将文件名中的空格替换为下划线
func ExampleRule(name string) string {
    return strings.ReplaceAll(name, " ", "_")
}

func main() {
    // 解析命令行参数
    dir := flag.String("dir", ".", "Directory to rename files in")
    flag.Parse()

    // 初始化 Renamer 实例
    renamer := NewRenamer(*dir, ExampleRule)

    // 执行重命名操作，并处理可能的错误
    if err := renamer.Rename(); err != nil {
        log.Fatalf("Error renaming files: %s
", err)
    }
}
