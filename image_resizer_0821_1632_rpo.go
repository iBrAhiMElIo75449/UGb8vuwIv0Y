// 代码生成时间: 2025-08-21 16:32:34
package main

import (
    "bufio"
    "encoding/json"
    "flag"
    "fmt"
    "image"
    "image/jpeg"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strconv"
    "strings"
)

// Config holds the configuration for the image resizer
type Config struct {
    Width  int    `json:"width"`
    Height int    `json:"height"`
    Input  string `json:"input"`
    Output string `json:"output"`
}

// resizeImage resizes an image to the specified dimensions
func resizeImage(img image.Image, width, height int) image.Image {
    dest := image.NewRGBA(image.Rect(0, 0, width, height))
    destRect := dest.Bounds()
    srcRect := img.Bounds()
    
    // Create a new image with the desired size
    draw.Draw(dest, destRect, image.NewUniform(color.White), image.Point{}, draw.Src)
    
    // Calculate the scaling factor to maintain aspect ratio
    srcW, srcH := srcRect.Dx(), srcRect.Dy()
    scaleX := float64(width) / float64(srcW)
    scaleY := float64(height) / float64(srcH)
    scale := min(scaleX, scaleY)
    srcRect = srcRect.Inset(src.Rect(srcW*scale - srcW, srcH*scale - srcH, srcW, srcH))
    
    // Draw the original image in the new image with the calculated scaling
    draw.Draw(dest, srcRect, img, image.Point{}, draw.Src)
    return dest
}

// min returns the minimum of two integers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// processImage resizes an image and saves it to the output directory
func processImage(config Config) error {
    imgFile, err := os.Open(config.Input)
    if err != nil {
        return err
    }
    defer imgFile.Close()
    
    img, _, err := image.Decode(imgFile)
    if err != nil {
        return err
    }
    
    resizedImg := resizeImage(img, config.Width, config.Height)
    
    outfile, err := os.Create(config.Output)
    if err != nil {
        return err
    }
    defer outfile.Close()
    
    err = jpeg.Encode(outfile, resizedImg, nil)
    if err != nil {
        return err
    }
    
    return nil
}

// parseConfig parses the configuration from a JSON file
func parseConfig(configPath string) (Config, error) {
    configFile, err := ioutil.ReadFile(configPath)
    if err != nil {
        return Config{}, err
    }
    var config Config
    err = json.Unmarshal(configFile, &config)
    if err != nil {
        return Config{}, err
    }
    return config, nil
}

func main() {
    configFile := flag.String("config", "config.json", "Path to the configuration JSON file")
    flag.Parse()
    
    config, err := parseConfig(*configFile)
    if err != nil {
        log.Fatalf("Error parsing config: %s", err)
    }
    
    if err := processImage(config); err != nil {
        log.Fatalf("Error resizing image: %s", err)
    }
    fmt.Println("Image resized successfully")
}
