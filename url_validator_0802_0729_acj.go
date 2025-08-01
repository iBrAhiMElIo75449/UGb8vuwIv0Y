// 代码生成时间: 2025-08-02 07:29:18
package main

import (
    "buffalo"
    "encoding/json"
    "net/http"
    "errors"
    "regexp"
)

// URLValidator represents the struct for URL validation
type URLValidator struct{}

// Validate checks if the provided URL is valid.
// It returns a JSON response with the validation result.
func (u *URLValidator) Validate(c buffalo.Context) error {
    // Retrieve the URL from the request query parameter
    urlStr := c.Request().URL.Query().Get("url")
    
    // Check if the URL is empty
    if urlStr == "" {
        return buffalo.NewErrorFrom(errEmptyURL, http.StatusBadRequest)
    }
    
    // Define the regular expression for URL validation
    urlRegex := regexp.MustCompile(`^(https?|ftp):\/\/[^\s/$.?#].[^\s]*$`)
    
    // Validate the URL using the regular expression
    if !urlRegex.MatchString(urlStr) {
        return buffalo.NewErrorFrom(errInvalidURL, http.StatusBadRequest)
    }
    
    // Return a JSON response with the validation result
    return c.Render(http.StatusOK, json.NewJSONRenderer().Render(c, json.Map{
        "valid": true,
        "message": "The URL is valid",
    }))
}

// Define custom error types for URL validation
var (
    errEmptyURL = errors.New("URL cannot be empty")
    errInvalidURL = errors.New("Invalid URL format")
)

func main() {
    app := buffalo.New(buffalo.Options{
        Env: "development",
    })
    
    // Define the route for URL validation
    app.GET("/validate-url", func(c buffalo.Context) error {
        return (&URLValidator{}).Validate(c)
    })
    
    // Run the application
    app.Serve()
}