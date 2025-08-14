// 代码生成时间: 2025-08-14 08:56:15
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "os"
    "buffalo"
    "github.com/gobuffalo/buffalo/generators"
    "log"
)

// HashCalculator represents our application's main struct
type HashCalculator struct {
    // you can embed other structs here if needed
}

// NewHashCalculator creates a new instance of HashCalculator
func NewHashCalculator() *HashCalculator {
    return &HashCalculator{}
}

// CalculateHash takes a string and calculates its SHA256 hash
func (h *HashCalculator) CalculateHash(input string) (string, error) {
    // Create a new hash.Hash
    hash := sha256.New()
    // Write the input to the hash
    if _, err := hash.Write([]byte(input)); err != nil {
        return "", fmt.Errorf("error writing to hash: %w", err)
    }
    // Return the hex-encoded hash
    return hex.EncodeToString(hash.Sum(nil)), nil
}

// main function to run our application
func main() {
    // Create a new buffalo app
    app := buffalo.App(
        // Generate the buffalo app
        func() *buffalo.App {
            a := buffalo.New(buffalo.Options{
                Name: "hash-calculator",
            })
            a.GET("/hash", func(c buffalo.Context) error {
                // Get the input string from the query parameter
                input := c.Param("input")
                if input == "" {
                    return c.Render(200, r.String("Please provide an input to calculate the hash"))
                }
                // Create a new HashCalculator instance
                calculator := NewHashCalculator()
                // Calculate the hash
                hash, err := calculator.CalculateHash(input)
                if err != nil {
                    return c.Render(500, r.String("Error calculating hash: "+err.Error()))
                }
                // Return the hash in the response
                return c.Render(200, r.JSON(map[string]string{"hash": hash}))
            })
            return a
        }(),
    )
    // Run the buffalo app
    log.Fatal(app.Serve(\os.Args))
}
