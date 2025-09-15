package api

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "strings"

    "github.com/gofiber/fiber/v2"
    "nadhi.dev/binaries/server/scrapper"
)

// Handler for GET /packages
func GetPackages(c *fiber.Ctx) error {
    data, err := ioutil.ReadFile("packages.json")
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not read packages.json"})
    }
    c.Type("json")
    return c.Send(data)
}

func SearchPackages(c *fiber.Ctx) error {
    query := strings.ToLower(c.Query("q"))
    if query == "" {
        return c.Status(400).JSON(fiber.Map{"error": "Query parameter 'q' is required"})
    }

    // Try to read local packages.json
    data, err := ioutil.ReadFile("packages.json")
    var pkgData struct {
        Packages map[string]interface{} `json:"packages"`
    }
   
    if err == nil {
        if err := json.Unmarshal(data, &pkgData); err == nil {
            result := make(map[string]interface{})
            for name, val := range pkgData.Packages {
                if strings.Contains(strings.ToLower(name), query) {
                    result[name] = val
                }
            }
            if len(result) > 0 {
                return c.JSON(fiber.Map{"packages": result})
            }
        }
    }

    // If not found locally, scrape pkg.go.dev
    html, err := scrapper.FetchSearchContent(query, 10)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch from pkg.go.dev"})
    }
    parsed, err := scrapper.ParseSearchContent(html)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not parse pkg.go.dev results"})
    }
    transformed := scrapper.TransformToClientFormat(parsed)

    // Update the local registar to make things faster next time.
    
    if len(transformed) > 0 {
        // Read existing packages.json if possible
        var allPkgs map[string]interface{}
        if data != nil {
            var existing struct {
                Packages map[string]interface{} `json:"packages"`
            }
            if err := json.Unmarshal(data, &existing); err == nil && existing.Packages != nil {
                allPkgs = existing.Packages
            }
        }
        if allPkgs == nil {
            allPkgs = make(map[string]interface{})
        }
        for k, v := range transformed {
            allPkgs[k] = v
        }
        out, _ := json.MarshalIndent(struct {
            Packages map[string]interface{} `json:"packages"`
        }{Packages: allPkgs}, "", "  ")
        _ = os.WriteFile("packages.json", out, 0644)
    }

    return c.JSON(fiber.Map{"packages": transformed})
}