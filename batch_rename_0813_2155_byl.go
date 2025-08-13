// 代码生成时间: 2025-08-13 21:55:17
package main

import (
    "os"
    "path/filepath"
    "strings"
    "fmt"
    "log"
)

// BatchRename renames files in a directory based on a pattern.
func BatchRename(dir string, pattern string, newPrefix string) error {
    // List all files in the directory.
    files, err := os.ReadDir(dir)
    if err != nil {
        return fmt.Errorf("error reading directory: %w", err)
    }

    for _, file := range files {
        if !file.IsDir() {
            // Construct the old file path.
            oldPath := filepath.Join(dir, file.Name())

            // Construct the new file name based on the pattern and prefix.
            extension := filepath.Ext(file.Name())
            newFileName := fmt.Sprintf("%s%s%s", newPrefix, pattern, extension)
            newPath := filepath.Join(dir, newFileName)

            // Rename the file.
            if err := os.Rename(oldPath, newPath); err != nil {
                return fmt.Errorf("error renaming file %s to %s: %w", oldPath, newPath, err)
            }
        }
    }
    return nil
}

func main() {
    // Specify the directory to rename files in.
    dir := "./files"
    // Specify the pattern to insert into the file names.
    pattern := "_renamed_"
    // Specify the new prefix for file names.
    newPrefix := "new_"

    // Perform the batch rename operation.
    if err := BatchRename(dir, pattern, newPrefix); err != nil {
        log.Fatalf("Failed to rename files: %s", err)
    } else {
        fmt.Println("Files have been renamed successfully.")
    }
}
