// 代码生成时间: 2025-08-12 15:15:08
package main

import (
    "os"
    "log"
    "buffalo"
    "github.com/markbates/buffalo/buffalo"
    "github.com/markbates/buffalo-pop/pop/pop"
)

// ConfigManager is the main application struct
type ConfigManager struct {
    *buffalo.App
    Pop *pop.Pop
}

// NewConfigManager creates a new instance of ConfigManager
func NewConfigManager(db *pop.Connection) *ConfigManager {
    return &ConfigManager{
        App: buffalo.New(buffalo.Options{
            Env:         os.Getenv("GO_ENV"),
            SessionStore: buffalosession.NewCookieStore([]byte("supersecret