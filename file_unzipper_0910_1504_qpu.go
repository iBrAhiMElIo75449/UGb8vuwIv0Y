// 代码生成时间: 2025-09-10 15:04:36
package main

import (
    "archive/zip"
    "bufio"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "github.com/gobuffalo/buffalo"
)

// FileUnzipper handler handles the file upload and unzips the uploaded zip file
func FileUnzipper(c buffalo.Context) error {
    // Get the uploaded file
    file, err := c.File("file")
    if err != nil {
        return err
    }
    defer file.Close()

    // The path to save the file temporarily
    filePath := filepath.Join(os.TempDir(), file.Filename)
    err = c.SaveFile("file", filePath)
    if err != nil {
        return err
    }

    // Create the destination directory
    destPath := filepath.Join(os.TempDir(), "unzipped")
    os.MkdirAll(destPath, os.ModePerm)

    // Open the zip file
    r, err := zip.OpenReader(filePath)
    if err != nil {
        return err
    }
    defer r.Close()

    // Iterate through the files in the zip
    for _, f := range r.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        // Create the file in the destination directory
        fPath := filepath.Join(destPath, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(fPath, os.ModePerm)
        } else {
            fDir := filepath.Dir(fPath)
            os.MkdirAll(fDir, os.ModePerm)
            outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer outFile.Close()

            // Copy the contents from the zip file to the new file
            _, err = io.Copy(outFile, rc)
            if err != nil {
                return err
            }
        }
    }

    // Return success message
    return c.Render(200, r.HTML("success.html", nil))
}

func main() {
    app := buffalo.Automatic()
    app.GET("/", FileUnzipper)
    app.Serve()
}
