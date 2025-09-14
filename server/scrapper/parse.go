package scrapper


import (
    "encoding/json"
   _ "fmt"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

// PackageResult represents a parsed package from the search result.
type PackageResult struct {
    Name        string `json:"name"`
    ImportPath  string `json:"import_path"`
    GoGet       string `json:"go_get"`
    Version     string `json:"version"`
    Description string `json:"description"`
    LastUpdated string `json:"last_updated"`
    License     string `json:"license"`
    ImportedBy  string `json:"imported_by"`
}

// ParseSearchContent parses the HTML content and returns a JSON array of packages.
func ParseSearchContent(htmlContent string) ([]PackageResult, error) {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
    if err != nil {
        return nil, err
    }

    var results []PackageResult

    doc.Find("div.SearchSnippet").Each(func(i int, s *goquery.Selection) {
        var pkg PackageResult

        // Name and ImportPath
        title := s.Find("h2 a[data-test-id='snippet-title']")
        pkg.Name = strings.TrimSpace(title.Contents().First().Text())
        href, _ := title.Attr("href")
        pkg.ImportPath = strings.Trim(strings.TrimPrefix(href, "/"), " ")
        pkg.GoGet = pkg.ImportPath // For most, go get path is the import path

        // Version and LastUpdated
        info := s.Find("div.SearchSnippet-infoLabel")
        pkg.Version = strings.TrimSpace(info.Find("strong").First().Text())
        pkg.LastUpdated = strings.TrimSpace(info.Find("span[data-test-id='snippet-published'] strong").Text())

        // Description
        pkg.Description = strings.TrimSpace(s.Find("p.SearchSnippet-synopsis").Text())

        // License
        pkg.License = strings.TrimSpace(info.Find("span[data-test-id='snippet-license'] a").Text())

        // ImportedBy
        pkg.ImportedBy = strings.TrimSpace(info.Find("a[aria-label='Go to Imported By'] strong").Text())

        results = append(results, pkg)
    })

    return results, nil
}

// ToJSON helper to convert results to JSON string (optional)
func ResultsToJSON(results []PackageResult) (string, error) {
    data, err := json.MarshalIndent(results, "", "  ")
    if err != nil {
        return "", err
    }
    return string(data), nil
}