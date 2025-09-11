// 代码生成时间: 2025-09-12 04:30:30
package main

import (
    "buffalo"
    "buffalo/render"
    "github.com/markbates/pkg/hash"
    "log"
)

// HashCalculatorHandler is a handler function that calculates the hash value of a given string.
// It uses the BUFFALO framework to handle HTTP requests.
func HashCalculatorHandler(c buffalo.Context) error {
    // Get the string parameter from the URL query or form data.
    str := c.Param("string")

    // Check if the string is empty.
    if str == "" {
        // Return a bad request error if the string is empty.
        return c.Error(400, "string parameter is required")
    }

    // Calculate the hash value using the hash function from the hash package.
    hashValue, err := hash.MD5.String(str)
    if err != nil {
        // Return an internal server error if the hash calculation fails.
        return c.Error(500, "failed to calculate hash")
    }

    // Render the hash value as JSON in the response.
    return c.Render(200, render.JSON(hashValue))
}

// main function to start the BUFFALO application.
func main() {
    // Create a new BUFFALO application.
    app := buffalo.Classic()

    // Define the main route for the hash calculator.
    app.GET("/hash/{string}", HashCalculatorHandler)

    // Start the BUFFALO application.
    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
