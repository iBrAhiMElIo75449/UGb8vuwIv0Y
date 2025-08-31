// 代码生成时间: 2025-08-31 08:46:51
package main

import (
    "archive/zip"
    "bufio"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
)

// Unzipper defines the structure for unzipping files
type Unzipper struct {
    // Path to the zip file to be extracted
    ZipFilePath string
    // Destination directory for the extracted files
    Destination string
}

// NewUnzipper initializes a new Unzipper instance
func NewUnzipper(zipFilePath, destination string) *Unzipper {
    return &Unzipper{ZipFilePath: zipFilePath, Destination: destination}
}

// Unzip extracts the contents of the zip file to the destination directory
func (u *Unzipper) Unzip() error {
    // Open the zip file
    reader, err := zip.OpenReader(u.ZipFilePath)
    if err != nil {
        return fmt.Errorf("failed to open zip file: %w", err)
    }
    defer reader.Close()

    // Iterate through the files in the zip
    for _, file := range reader.File {
        filePath := filepath.Join(u.Destination, file.Name)
        if file.FileInfo().IsDir() {
            // Create the directory structure
            if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
                return fmt.Errorf("failed to create directory: %w", err)
            }
            continue
        }

        // Create the file structure
        if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
            return fmt.Errorf("failed to create parent directory: %w", err)
        }

        // Open the file for writing
        fileWriter, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
        if err != nil {
            return fmt.Errorf("failed to open file for writing: %w", err)
        }
        defer fileWriter.Close()

        // Copy the contents of the file from the zip to the destination
        fileReader, err := file.Open()
        if err != nil {
            return fmt.Errorf("failed to open file inside zip: %w", err)
        }
        defer fileReader.Close()

        _, err = io.Copy(fileWriter, fileReader)
        if err != nil {
            return fmt.Errorf("failed to copy file: %w", err)
        }
    }
    return nil
}

func main() {
    zipPath := flag.String("zip", "", "Path to the zip file")
    destPath := flag.String("dest", "", "Destination directory for extraction")
    flag.Parse()

    if *zipPath == "" || *destPath == "" {
        log.Fatal("Both zip and destination paths must be provided")
    }

    unzipper := NewUnzipper(*zipPath, *destPath)
    if err := unzipper.Unzip(); err != nil {
        log.Fatalf("Failed to unzip file: %s", err)
    }
    fmt.Println("Unzip successful")
}
