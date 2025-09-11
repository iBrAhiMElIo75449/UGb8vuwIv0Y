// 代码生成时间: 2025-09-11 21:08:05
package main

import (
    "bufio"
    "errors"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// BatchRenamer 结构体封装了文件重命名操作的逻辑
type BatchRenamer struct {
    srcPath string // 源文件路径
    destPath string // 目标文件路径
    pattern string // 重命名模式
}

// NewBatchRenamer 创建并返回一个BatchRenamer实例
func NewBatchRenamer(srcPath, destPath, pattern string) *BatchRenamer {
    return &BatchRenamer{
        srcPath: srcPath,
        destPath: destPath,
        pattern: pattern,
    }
}

// RenameFiles 执行批量文件重命名操作
func (br *BatchRenamer) RenameFiles() error {
    srcFiles, err := os.ReadDir(br.srcPath)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range srcFiles {
        if !file.IsDir() {
            srcFilePath := filepath.Join(br.srcPath, file.Name())
            destFilePath := filepath.Join(br.destPath, br.renameFile(file.Name()))
            
            if err := os.Rename(srcFilePath, destFilePath); err != nil {
                return fmt.Errorf("failed to rename file %s to %s: %w", srcFilePath, destFilePath, err)
            }
        }
    }
    return nil
}

// renameFile 根据提供的模式对文件名进行重命名
func (br *BatchRenamer) renameFile(fileName string) string {
    return fmt.Sprintf(br.pattern, strings.ReplaceAll(filepath.Ext(fileName), ".", ""))
}

func main() {
    srcPath := "./src"
    destPath := "./dest"
    pattern := "new_%d.%s" // 例如：new_1.txt

    br := NewBatchRenamer(srcPath, destPath, pattern)
    if err := br.RenameFiles(); err != nil {
        log.Fatalf("Error renaming files: %s", err)
    }
    fmt.Println("Files have been renamed successfully.")
}
