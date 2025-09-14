package nadhi

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// PackageData represents the structure of the packages.json response
type PackageData struct {
    Packages map[string]interface{} `json:"packages"`
}

// FetchPackages fetches all packages from the server API
func FetchPackages(apiURL string) (*PackageData, error) {
    resp, err := http.Get(apiURL + "/packages")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var data PackageData
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, err
    }
    return &data, nil
}

// SearchPackages fetches packages matching the query from the server API
func SearchPackages(apiURL, query string) (*PackageData, error) {
    url := fmt.Sprintf("%s/packages/search?q=%s", apiURL, query)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var data PackageData
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, err
    }
    return &data, nil
}