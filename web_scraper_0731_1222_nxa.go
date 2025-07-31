// 代码生成时间: 2025-07-31 12:22:57
package main

import (
    "buffalo"
    "github.com/PuerkitoBio/goquery"
    "net/http"
    "log"
)

// WebScraper is the main application struct
type WebScraper struct {
    *buffalo.App
}

// NewWebScraper creates a new instance of the WebScraper
func NewWebScraper(a *buffalo.App) *WebScraper {
    return &WebScraper{App: a}
}

// ScrapeHandler handles the scraping of web pages
func (ws *WebScraper) ScrapeHandler(c buffalo.Context) error {
    var err error
    defer func() {
        if re := recover(); re != nil {
            log.Printf("Recovered in ScrapeHandler: %v", re)
        }
    }()

    // Fetch the URL from the query parameters
    url := c.Request().URL.Query().Get("url")
    if url == "" {
        return buffalo.NewErrorPage(400, "URL parameter is required")
    }

    // Send an HTTP GET request to the target URL
    resp, err := http.Get(url)
    if err != nil {
        return buffalo.NewErrorPage(404, "Failed to fetch the URL")
    }
    defer resp.Body.Close()

    // Parse the HTML content using goquery
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return buffalo.NewErrorPage(500, "Failed to parse HTML")
    }

    // Use goquery to extract content from the HTML document
    // This is a simple example, you can extend this to fit your needs
    title := doc.Find("title").Text()
    body := doc.Find("body").Text()

    // Return the scraped content as JSON
    c.Set("title", title)
    c.Set("body", body)
    return c.Render(200, buffalo.R.JSON("{\"title\":\"" + title + "\",\"body\":\"" + body + "\"}"))
}

// main is the entry point of the application
func main() {
    app := buffalo.New(buffalo.Options{
        Addr: "127.0.0.1:3000",
    })
    defer app.Close()

    ws := NewWebScraper(app)

    // Define the route for the scraping handler
    app.GET("/scrap", ws.ScrapeHandler)

    // Start the application
    log.Fatal(app.Start())
}
