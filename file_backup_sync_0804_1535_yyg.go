// 代码生成时间: 2025-08-04 15:35:47
package main

import (
    "log"
    "os"
    "path/filepath"
    "time"
    "github.com/buffalo/buffalo"
)

// BackupConfig 配置备份参数
type BackupConfig struct {
    SourceDir  string    `json:"source_dir"`  // 源目录
    BackupDir  string    `json:"backup_dir"`  // 备份目录
    SyncTarget string    `json:"sync_target"` // 同步目标目录
    Interval   time.Duration // 同步间隔时间
}

// backup 备份文件
func backup(cfg *BackupConfig) error {
    err := filepath.WalkDir(cfg.SourceDir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil
        }
        relPath, err := filepath.Rel(cfg.SourceDir, path)
        if err != nil {
            return err
        }
        backupPath := filepath.Join(cfg.BackupDir, relPath)
        if _, err := os.Stat(backupPath); os.IsNotExist(err) {
            if err := os.MkdirAll(filepath.Dir(backupPath), 0755); err != nil {
                return err
            }
        }
        srcFile, err := os.Open(path)
        if err != nil {
            return err
        }
        defer srcFile.Close()
        dstFile, err := os.Create(backupPath)
        if err != nil {
            return err
        }
        defer dstFile.Close()
        if _, err := io.Copy(dstFile, srcFile); err != nil {
            return err
        }
        return nil
    })
    return err
}

// sync 同步目录
func sync(cfg *BackupConfig) error {
    err := filepath.WalkDir(cfg.SourceDir, func(path string, d os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        relPath, err := filepath.Rel(cfg.SourceDir, path)
        if err != nil {
            return err
        }
        targetPath := filepath.Join(cfg.SyncTarget, relPath)
        if d.IsDir() {
            if _, err := os.Stat(targetPath); os.IsNotExist(err) {
                if err := os.MkdirAll(targetPath, 0755); err != nil {
                    return err
                }
            }
            return nil
        }
        if _, err := os.Stat(targetPath); os.IsNotExist(err) {
            srcFile, err := os.Open(path)
            if err != nil {
                return err
            }
            defer srcFile.Close()
            dstFile, err := os.Create(targetPath)
            if err !=
            }
        }
    })
    return err
}

func main() {
    cfg := &BackupConfig{
        SourceDir:  "/path/to/source",
        BackupDir:  "/path/to/backup",
        SyncTarget: "/path/to/sync",
        Interval:   5 * time.Minute,
    }

    go func() {
        for {
            if err := backup(cfg); err != nil {
                log.Printf("Backup error: %v", err)
            }
            time.Sleep(cfg.Interval)
        }
    }()

    go func() {
        for {
            if err := sync(cfg); err != nil {
                log.Printf("Sync error: %v", err)
            }
            time.Sleep(cfg.Interval)
        }
    }()

    // 启动BUFFALO应用
    app := buffalo.Buffalo(buffalo.Options{
        ServiceProviders: []buffalo.ServiceProvider{
            buffalo.FileServer{"public/"},
        },
    })
    app.Serve()
}