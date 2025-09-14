package scrapper

import (
    "fmt"
    "io"
    "net/http"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

// FetchSearchContent fetches and returns only the HTML content inside the div with class "go-Content SearchResults".
func FetchSearchContent(query string, limit int) (string, error) {
    url := fmt.Sprintf("https://pkg.go.dev/search?limit=%d&m=package&q=%s", limit, query)
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to fetch: %s", resp.Status)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
    if err != nil {
        return "", err
    }

    var content string
    doc.Find("div.go-Content.SearchResults").Each(func(i int, s *goquery.Selection) {
        html, err := s.Html()
        if err == nil {
            content = html
        }
    })

    if content == "" {
        return "", fmt.Errorf("content div not found")
    }

    return content, nil
}